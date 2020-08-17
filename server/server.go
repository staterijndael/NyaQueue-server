package server

import (
	"nyaqueue-server/adapter"
	"nyaqueue-server/queue"

	"github.com/Oringik/otty"
	"go.uber.org/zap"
)

// Server ...
type Server struct {
	ottyStruct  *otty.Otty
	queueList   map[uint]*queue.Queue
	adapterList map[uint]*adapter.Adapter
	RoutingID   uint
	logger      *zap.Logger
}

// NewServer ...
func NewServer(logger *zap.Logger) *Server {
	return &Server{
		ottyStruct:  otty.New(),
		logger:      logger,
		queueList:   make(map[uint]*queue.Queue),
		adapterList: make(map[uint]*adapter.Adapter),
	}
}

// ListenDataChannel ...
func (server *Server) ListenDataChannel(dataChannel <-chan []byte) {
	for {
		data := <-dataChannel
		server.ottyStruct.ParseOtty(data)
		server.ottyStruct.ResolveEndpoint(server.ottyStruct.Route().GetValue(), server.ottyStruct.Data().GetValue())
	}
}

// CreateEndpoints ...
func (server *Server) CreateEndpoints() {

	server.ottyStruct.CreateEndpoint("createQueue", func(data []byte) {
		err := server.createQueue(data)
		if err != nil {
			server.logger.Panic(err.Error(),
				zap.String("endpointName", "createQueue"),
			)
		}

		return
	})

}
