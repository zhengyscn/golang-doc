package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

func main() {
	/*
		证书的有效期
		接入层(nginx)
	*/
	conn, err := tls.Dial("tcp", "github.com:443", nil)
	if err != nil {
		log.Fatal(err)
	}
	conn_state := conn.ConnectionState()
	fmt.Println(conn_state.Version)
	fmt.Println(conn_state.ServerName)
	fmt.Println(conn_state.PeerCertificates)

	for _, pc := range conn_state.PeerCertificates {
		fmt.Printf("isca: %s\n", pc.IsCA)
		fmt.Printf("NotBefore: %s, NotAfter %s\n", pc.NotBefore, pc.NotAfter)
	}
}
