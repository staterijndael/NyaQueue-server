package main

import (
	"log"
	"nyaqueue-server/network"
)

func main() {

	conn, err := network.StartTCPServer("127.0.0.1:12000")
	if err != nil {
		conn.Close()
		log.Fatal(err)
	}

	listenChannel := make(chan []byte)

	network.ListenConn()

}
