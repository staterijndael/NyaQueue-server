package adapter

import (
	"encoding/json"
	"nyaqueue-server/queue"
)

type adapterType int

const (
	// FANOUT is a type for adapter to send message for all queues in adapter list
	FANOUT adapterType = iota
	// DIRECT is a type for adapter to send message for only one queue which defined in type field
	DIRECT
)

// Adapter ...
type Adapter struct {
	Name      string
	QueueList map[uint]*queue.Queue
	Type      adapterType
}

// NewAdapter ...
func NewAdapter() *Adapter {
	return &Adapter{
		Name:      "",
		QueueList: map[uint]*queue.Queue{},
		Type:      DIRECT,
	}
}

// ResolveData ...
func (ad *Adapter) ResolveData(data string) error {
	type resolveInfoStruct struct {
		RoutingID uint
		Data      interface{}
	}

	resolveInfoStructInstance := &resolveInfoStruct{}

	if err := json.Unmarshal([]byte(data), resolveInfoStructInstance); err != nil {
		return err
	}

	if ad.Type == DIRECT {
		neededQueue := ad.QueueList[resolveInfoStructInstance.RoutingID]

		neededQueue.WriteInto(resolveInfoStructInstance.Data)
	} else if ad.Type == FANOUT {
		neededQueues := ad.QueueList

		for _, queue := range neededQueues {
			queue.WriteInto(resolveInfoStructInstance.Data)
		}
	}

	return nil

}
