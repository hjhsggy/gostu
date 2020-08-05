package main

import (
	"encoding/binary"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Println(err)
		return
	}

	var headSize int
	var headByte = make([]byte, 2)
	msg := "hello golang"
	content := []byte(msg)
	headSize = len(content)

	binary.BigEndian.PutUint16(headByte, uint16(headSize))
	conn.Write(headByte)
	conn.Write(content)

	select {}
}
