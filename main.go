package main

import (
	"log"
	"nyaqueue-server/network"
	"nyaqueue-server/server"
	"runtime"
	"time"

	"go.uber.org/zap"
)

func main() {
	startTime := time.Now().UnixNano()

	runtime.GOMAXPROCS(runtime.NumCPU())

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer logger.Sync()

	ln, err := network.StartTCPServer("127.0.0.1:12000")
	if err != nil {
		log.Fatal(err.Error())
	}

	listenChannel := make(chan []byte)

	go network.ListenConn(ln, logger, listenChannel)

	server := server.NewServer(logger)

	server.CreateEndpoints()

	logger.Info("Server successfully started",
		zap.Int64("Duration", time.Now().UnixNano()-startTime),
	)

	server.ListenDataChannel(listenChannel)

}
