package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gomodule/redigo/redis"
)

//构造一个链接函数，如果没有密码，passwd为空字符串
func redisConn(host, port, passwd string) (redis.Conn, error) {
	c, err := redis.Dial("tcp",
		host+":"+port,
		redis.DialConnectTimeout(5*time.Second),
		redis.DialReadTimeout(1*time.Second),
		redis.DialWriteTimeout(1*time.Second),
		redis.DialPassword(passwd),
		redis.DialKeepAlive(1*time.Second),
	)
	return c, err
}

func NewRedisPool(host, port, passwd string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:         5,  //定义redis连接池中最大的空闲链接为3
		MaxActive:       18, //在给定时间已分配的最大连接数(限制并发数)
		IdleTimeout:     240 * time.Second,
		MaxConnLifetime: 300 * time.Second,
		Dial:            func() (redis.Conn, error) { return redisConn(host, port, passwd) },
	}
}

func worker_production(c redis.Conn) {
	var i int = 0
	for {
		c.Send("RPUSH", "rq", fmt.Sprintf("zhengysn-%d", i))
		if err := c.Flush(); err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
			continue
		}
		if reply, err := c.Receive(); err != nil {
			fmt.Printf("Receive err:%v\n", err)
			time.Sleep(time.Second)
			continue
		} else {
			_ = reply
		}
		time.Sleep(time.Second)
		i++
	}

}

func worker_consumer(c redis.Conn) {
	for {
		c.Send("RPOP", "rq")
		if err := c.Flush(); err != nil {
			fmt.Printf("Flush err:%v\n", err)
			time.Sleep(time.Second)
			continue
		}
		if reply, err := c.Receive(); err != nil {
			fmt.Println("Receive err:%v\n", err)
			time.Sleep(time.Second)
			continue
		} else {
			v, ok := reply.([]byte)
			if !ok {
				fmt.Println("type asserts false.")
				time.Sleep(time.Second)
				continue
			}
			fmt.Printf("Receive value:%v\n", string(v))
		}

		time.Sleep(time.Second)
	}
}

func main() {
	//var lock sync.Mutex

	p := NewRedisPool("0.0.0.0", "6379", "")
	defer p.Close()

	// 从池里获取连接
	conn := p.Get()

	// 用完后将连接放回连接池
	defer conn.Close()
	runtime.GOMAXPROCS(runtime.NumCPU())
	go worker_production(conn)
	go worker_consumer(conn)

	time.Sleep(time.Second * 30)
	fmt.Println("Over.")
}
