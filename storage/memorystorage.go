package storage

import "sync"

// TransportData ...
type TransportData struct {
	Info      interface{}
	RoutingID uint
}

// NewTransportData ...
func NewTransportData() TransportData {
	return TransportData{
		Info: nil,
	}
}

// MemoryStorage ...
type MemoryStorage struct {
	Data     []TransportData
	BusyData map[uint]TransportData
	mx       sync.RWMutex
}

// NewMemStore ...
func NewMemStore() *MemoryStorage {
	return &MemoryStorage{
		Data: []TransportData{},
	}
}

// Write ...
func (storage *MemoryStorage) Write(data interface{}) {
	storage.mx.Lock()

	transportData := NewTransportData()
	transportData.Info = data

	storage.Data = append(storage.Data, transportData)
	storage.mx.Unlock()
}

// ReadLast ...
func (storage *MemoryStorage) ReadLast() TransportData {
	storage.mx.RLock()
	lastRecord := storage.Data[len(storage.Data)-1]
	storage.mx.RUnlock()

	return lastRecord
}
