package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/zhengyscn/golang-doc/tcp/utils"
)

func CreateTcpServer(t *utils.ServerInfo) {
	addr := fmt.Sprintf("%s:%d", t.Host, t.Port)
	log.Println("Listening......")
	listener, err := net.Listen(t.Proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		// Wait for a connection.
		log.Println("Accept......")
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("come a new connection from %s\n", conn.RemoteAddr().String())

		var wg sync.WaitGroup
		wg.Add(1)
		go handleConn(conn, &wg)
	}
}

func handleConn(conn net.Conn, wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("Closing connection...")
		wg.Done()
		conn.Close()
	}()

	for {
		var buf = make([]byte, 1024)
		// Read
		n, err := conn.Read(buf) // ?
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(n, string(buf[:n]))
	}

}

func render(ch chan []byte) {
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		}
	}
}

func unPacket(buffer []byte, ch chan []byte) []byte {
	length := len(buffer)
	var i int
	for i = 0; i < length; i++ {
		if length < i+4 {
			break
		}
		msgLength := utils.ByteToInt(buffer[i : i+4])
	}
	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]

}

func handleConnV1(conn net.Conn, wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("Closing connection...")
		wg.Done()
		conn.Close()
	}()

	for {
		// Read
		content, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Print(err)
			return
		}
		fmt.Println(content)

		// Write
		n, err := conn.Write([]byte(content))
		if err != nil {
			log.Print(err)
			return
		}
		fmt.Printf("write %d data\n", n)
	}

}

func main() {
	serverinfo, err := utils.NewServer("127.0.0.1", "tcp", 2000)
	if err != nil {
		log.Fatal(err)
	}
	CreateTcpServer(serverinfo)

}
