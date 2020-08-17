package network

import (
	"io"
	"net"

	"go.uber.org/zap"
)

// StartTCPServer ...
func StartTCPServer(addr string) (net.Listener, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return ln, nil

}

// ListenConn ...
func ListenConn(ln net.Listener, logger *zap.Logger, sendDataChannel chan<- []byte) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Error(err.Error())
		}

		var readData []byte
		for {

			localReadData := make([]byte, 1024)

			_, err = conn.Read(localReadData)
			if err != nil {
				if err == io.EOF {
					readData = append(readData, localReadData...)
					sendDataChannel <- readData
					readData = []byte{}
					break
				}
				logger.Error(err.Error())
			}
			readData = append(readData, localReadData...)
		}
	}

}
