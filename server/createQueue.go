package server

import (
	"encoding/json"
	"nyaqueue-server/queue"
)

// RegisterQueueRequest ...
type RegisterQueueRequest struct {
	Name          string `json:"name"`
	isAckAdapter  bool   `json:"ack_adapter"`
	isAckProducer bool   `json:"ack_producer"`
	BindingID     uint   `json:"binding_id"`
}

// RegisterQueue ...
func createQueue(data []byte, queueList map[uint]*queue.Queue) error {
	request := &RegisterQueueRequest{}

	err := json.Unmarshal(data, request)
	if err != nil {
		return err
	}

	newQueueToRegister := queue.NewQueue(request.Name)

	queueList[uint(len(queueList))] = newQueueToRegister

	return nil

}
