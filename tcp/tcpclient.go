package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/zhengyscn/golang-doc/tcp/utils"
)

func ConnectTcpServer(t *utils.ServerInfo) {
	//var data string = "{\"username\" : \"zhengyscn\", \"password\" : \"123456\"}"
	addr := fmt.Sprintf("%s:%d", t.Host, t.Port)
	conn, err := net.DialTimeout(t.Proto, addr, time.Second*5)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	sender(conn)
	/*
		conn.SetDeadline(time.Now().Add(time.Second * 5))
		for {
			io.Copy(conn, os.Stdin)
		}
	*/

}

func Packet(msg []byte) []byte {
	s := make([]byte, 0)
	return append(append(s, utils.IntToByte(len(msg))...), msg...)
}

func sender(conn net.Conn) {
	for i := 0; i < 300; i++ {
		data := fmt.Sprintf("{\"ID\":%d, \"Name\":\"zhengyscn-%d\"}", i, i)
		// Write
		//n, err := conn.Write([]byte(data)) //TODO
		n, err := conn.Write(Packet([]byte(data))) //TODO
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Printf("write %d data\n", n)
	}
}

func sender1(conn net.Conn) {
	for i := 0; i < 10; i++ {
		data := fmt.Sprintf("{\"ID\":%d, \"Name\":\"zhengyscn-%d\"}", i, i)
		fmt.Println(data)
		// Write
		n, err := conn.Write([]byte(data))
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Printf("write %d data\n", n)

		// Read
		content, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(content)
	}
}

func main() {
	serverinfo, err := utils.NewServer("127.0.0.1", "tcp", 2000)
	if err != nil {
		log.Fatal(err)
	}
	ConnectTcpServer(serverinfo)
}
