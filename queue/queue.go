package queue

import (
	"nyaqueue-server/storage"
	"sync"
)

// Queue ...
type Queue struct {
	Name          string
	isAckAdapter  bool
	isAckProducer bool
	BindingID     uint
	Storage       *storage.MemoryStorage
	sync.RWMutex
}

// NewQueue ...
func NewQueue(name string) *Queue {
	return &Queue{
		Name:    "",
		Storage: storage.NewMemStore(),
	}
}

// WriteInto ...
func (q *Queue) WriteInto(data interface{}) {
	q.Lock()
	defer q.Unlock()

	q.Storage.Write(data)

}

// ReadLastFrom ...
func (q *Queue) ReadLastFrom() storage.TransportData {
	q.RLock()
	defer q.RUnlock()

	data := q.Storage.ReadLast()

	return data
}
