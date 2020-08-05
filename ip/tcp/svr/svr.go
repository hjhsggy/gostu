package svr

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

const (
	BYTES_SIZE uint16 = 1024
	HEAD_SIZE  int    = 2
)

// StartServer start server
func StartServer() {

	// listen
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Println(err)
		return
	}

	// accept
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		// deal conn
		go doConn2(conn)
	}

}

func doConn(conn net.Conn) {

	var (
		buffer           = bytes.NewBuffer(make([]byte, 0, BYTES_SIZE))
		bytes            = make([]byte, BYTES_SIZE)
		isHeader    bool = true
		contentSize int
		header      = make([]byte, HEAD_SIZE)
		content     = make([]byte, BYTES_SIZE)
	)

	// parse conn
	for {
		// read lens
		readLen, err := conn.Read(bytes)
		if err != nil {
			log.Println(err)
			return
		}
		// write to buffer
		_, err = buffer.Write(bytes[:readLen])
		if err != nil {
			log.Println(err)
			return
		}

		for {
			// head
			if isHeader {
				if buffer.Len() >= HEAD_SIZE {
					_, err := buffer.Read(header)
					if err != nil {
						log.Println(err)
						return
					}
					contentSize = int(binary.BigEndian.Uint16(header))
					isHeader = false
				}
			} else {
				break
			}

			// content
			if !isHeader {
				if buffer.Len() >= contentSize {
					_, err := buffer.Read(content[:contentSize])
					if err != nil {
						log.Println(err)
						return
					}
					fmt.Println(string(content[:contentSize]))
					isHeader = true
				}
			} else {
				break
			}
		}
	}
}

func doConn2(conn net.Conn) {
	var (
		buf         = newBuffer(conn, 16)
		headBuf     []byte
		contentBuf  []byte
		contentSize int
	)
	for {
		_, err := buf.readFromReader()
		if err != nil {
			log.Println(err)
			return
		}
		for {
			headBuf, err = buf.Seek(HEAD_SIZE)
			if err != nil {
				break
			}
			contentSize = int(binary.BigEndian.Uint16(headBuf))
			if buf.Len() >= contentSize-HEAD_SIZE {
				contentBuf = buf.read(HEAD_SIZE, contentSize)
				fmt.Println(string(contentBuf))
				continue
			}
			break
		}
	}
}

type buffer struct {
	reader io.Reader
	buf    []byte
	start  int
	end    int
}

func newBuffer(reader io.Reader, len int) buffer {
	buf := make([]byte, len)
	return buffer{reader, buf, 0, 0}
}

func (b *buffer) read(offset, n int) []byte {
	b.start += offset
	buf := b.buf[b.start : b.start+n]
	b.start += n
	return buf
}

func (b *buffer) Seek(n int) ([]byte, error) {
	if b.end-b.start >= n {
		buf := b.buf[b.start : b.start+n]
		return buf, nil
	}
	return nil, errors.New("")
}

func (b *buffer) grow() {
	if b.start == 0 {
		return
	}
	copy(b.buf, b.buf[b.start:b.end])
	b.end -= b.start
	b.start = 0
}

func (b *buffer) Len() int {
	return b.end - b.start
}

func (b *buffer) readFromReader() (int, error) {
	b.grow()
	n, err := b.reader.Read(b.buf[b.end:])
	if err != nil {
		return n, err
	}
	b.end += n
	return n, nil
}
