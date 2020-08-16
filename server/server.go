package server

import (
	"nyaqueue-server/queue"

	"github.com/Oringik/otty"
	"go.uber.org/zap"
)

// Server ...
type Server struct {
	ottyStruct *otty.Otty
	queueList  map[uint]*queue.Queue
	logger     *zap.Logger
}

// New ...
func New(logger *zap.Logger) *Server {
	return &Server{
		ottyStruct: otty.New(),
		logger:     logger,
	}
}

// ListenDataChannel ...
func (server *Server) ListenDataChannel(dataChannel chan []byte) {
	for {
		data := <-dataChannel
		server.ottyStruct = otty.ParseOtty(data)
		server.ottyStruct.ResolveEndpoint(server.ottyStruct.Route().GetValue(), server.ottyStruct.Data().GetValue())
	}
}

// CreateEndpoints ...
func (server *Server) CreateEndpoints() {

	server.ottyStruct.CreateEndpoint("createQueue", func(data []byte) interface{} {
		err := createQueue(data, server.queueList)
		if err != nil {
			server.logger.Panic(creatingEndpointError,
				zap.String("endpointName", "createQueue"),
			)
			return err
		}

		return nil
	})

}
