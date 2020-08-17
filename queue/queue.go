package queue

import (
	"nyaqueue-server/storage"
	"sync"
)

// Queue ...
type Queue struct {
	ID            uint
	Name          string
	isAckAdapter  bool
	isAckProducer bool
	BindingID     uint
	Storage       *storage.MemoryStorage
	sync.RWMutex
}

// NewQueue ...
func NewQueue(id uint, name string, isAckAd bool, isAckPr bool, bindingID uint) *Queue {
	return &Queue{
		ID:            id,
		Name:          name,
		isAckAdapter:  isAckAd,
		isAckProducer: isAckPr,
		BindingID:     bindingID,
		Storage:       storage.NewMemStore(),
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
