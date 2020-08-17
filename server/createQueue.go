package server

import (
	"encoding/json"
	"nyaqueue-server/adapter"
	"nyaqueue-server/queue"
)

// RegisterQueueRequest ...
type RegisterQueueRequest struct {
	Name          string `json:"name"`
	IsAckAdapter  bool   `json:"ack_adapter"`
	IsAckProducer bool   `json:"ack_producer"`
}

func (server *Server) createQueue(data []byte) error {
	request := &RegisterQueueRequest{}

	err := json.Unmarshal(data, request)
	if err != nil {
		return err
	}

	// Register new Queue

	queueID := uint(len(server.queueList))

	newQueueToRegister := queue.NewQueue(queueID, request.Name, request.IsAckAdapter, request.IsAckProducer, uint(len(server.queueList)))

	server.queueList[queueID] = newQueueToRegister

	// Register new adapter for attach to queue

	adapterID := uint(len(server.adapterList))

	newAdapterToRegister := adapter.NewAdapter(adapterID, request.Name+"Adapter")

	server.adapterList[adapterID] = newAdapterToRegister

	newAdapterToRegister.QueueList[queueID] = newQueueToRegister

	return nil

}
