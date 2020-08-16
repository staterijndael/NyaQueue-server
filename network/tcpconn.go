package network

import (
	"io"
	"log"
	"net"
)

// StartTCPServer ...
func StartTCPServer(addr string) (net.Conn, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	conn, err := ln.Accept()
	if err != nil {
		return nil, err
	}
	return conn, nil

}

// ListenConn ...
func ListenConn(conn net.Conn, sendDataChannel chan []byte) {
	var readData []byte
	for {
		localReadData := make([]byte, 1024)
		_, err := conn.Read(localReadData)
		if err != nil {
			if err == io.EOF {
				readData = append(readData, localReadData...)
				sendDataChannel <- readData
				readData = []byte{}
			}
			log.Fatal(err)
		}
		readData = append(readData, localReadData...)
	}
}
