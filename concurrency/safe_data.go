package concurrency

import "sync"

type Data struct {
	Attr1 string
	Attr2 string
	Attr3 string
	Attr4 int64
	Attr5 float64
}

type SafeData interface {
	GetData() *Data
	SetData(newData *Data)
}

type SafeDataNonBlockingRead struct {
	data     *Data
	dataLock sync.RWMutex
}

func NewSafeDataNonBlockingRead(d *Data) *SafeDataNonBlockingRead {
	return &SafeDataNonBlockingRead{
		data: d,
	}
}

func (d *SafeDataNonBlockingRead) GetData() *Data {
	d.dataLock.RLock()
	defer d.dataLock.RUnlock()
	return d.data
}

func (d *SafeDataNonBlockingRead) SetData(newData *Data) {
	d.dataLock.Lock()
	defer d.dataLock.Unlock()
	d.data = newData
}

type SafeDataBlockingRead struct {
	data     *Data
	dataLock sync.Mutex
}

func NewSafeDataBlockingRead(d *Data) *SafeDataBlockingRead {
	return &SafeDataBlockingRead{
		data: d,
	}
}

func (d *SafeDataBlockingRead) GetData() *Data {
	d.dataLock.Lock()
	defer d.dataLock.Unlock()
	return d.data
}

func (d *SafeDataBlockingRead) SetData(newData *Data) {
	d.dataLock.Lock()
	defer d.dataLock.Unlock()
	d.data = newData
}
