package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

type ServerInfo struct {
	Host  string
	Port  int
	Proto string
}

func NewServer(host, proto string, port int) (*ServerInfo, error) {
	if len(host) == 0 || len(proto) == 0 || port <= 0 {
		return nil, fmt.Errorf("host, proto or port error")
	}
	return &ServerInfo{
		Host:  host,
		Port:  port,
		Proto: proto,
	}, nil
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func IntToByte(n int) []byte {
	x := int32(n)
	byteBuf := bytes.NewBuffer([]byte{})
	binary.Write(byteBuf, binary.BigEndian, x)
	return byteBuf.Bytes()
}

func ByteToInt(b []byte) int {
	var (
		byteBuf = bytes.NewBuffer(b)
		x       int32
	)
	binary.Read(byteBuf, binary.BigEndian, &x)
	return int(x)
}
