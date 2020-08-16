package main

import (
	"fmt"
	"log"
	"nyaqueue-server/network"
	"nyaqueue-server/server"
	"runtime"
	"time"

	"go.uber.org/zap"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	startTime := time.Now().UnixNano()

	ln, err := network.StartTCPServer("127.0.0.1:12000")
	if err != nil {
		log.Fatal(err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer logger.Sync()

	listenChannel := make(chan []byte)

	go network.ListenConn(ln, logger, listenChannel)

	server := server.NewServer(logger)

	logger.Info("Server successfully started",
		zap.Int64("Duration", time.Now().UnixNano()-startTime),
	)

	server.ListenDataChannel(listenChannel)

}
