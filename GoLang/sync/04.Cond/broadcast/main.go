package main

import (
	"time"
	"fmt"
	"io"
	"bytes"
	"sync"
)

type DataBucket struct {
	buffer *bytes.Buffer
	mutex *sync.RWMutex
	cond *sync.Cond
}

func NewDataBucket() *DataBucket {
	buf := make([]byte, 0)
	db := &DataBucket{
		buffer : bytes.NewBuffer(buf),
		mutex : new(sync.RWMutex),
	}
	db.cond = sync.NewCond(db.mutex.RLocker())
	return db
}

func (db *DataBucket) Reader(i int) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	var data []byte
	var d byte
	var err error

	for {
		if d, err = db.buffer.ReadByte(); err != nil {
			if err == io.EOF {
				if string(data) != "" {
					fmt.Printf("Reader-%d: %s\n",i,data)
				}
				db.cond.Wait()
				data = data[:0]
				continue
			}
		}
		data = append(data,d)
	}
}

func (db *DataBucket) Writer(data []byte) (int, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	n, err := db.buffer.Write(data)
	db.cond.Broadcast()
	return n,err
}

func main()  {
	db := NewDataBucket()
	for i:=1; i<3; i++{
		go db.Reader(i)
	}
	for j:=0; j<10; j++ {
		go func (i int)  {
			d := fmt.Sprintf("data-%d",i)
			db.Writer([]byte(d))
		}(j)
	}
	time.Sleep(100*time.Millisecond)
}