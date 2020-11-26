package main

import (
	"time"
	"fmt"
	"io"
	"bytes"
	"sync"
)
//how to use Buffer
//A Buffer is a variable-sized buffer of bytes with Read and Write methods.
type DataBucket struct {
	buffer *bytes.Buffer
	mutex *sync.RWMutex
	cond *sync.Cond
}

func NewDataBucket() *DataBucket  {
	buf := make([]byte, 0) //宣告一個變數buf，並初始化為一個內容類型為byte的slice
	db := &DataBucket{
		buffer: bytes.NewBuffer(buf), //以buf(一個slice)初始化buffer
		mutex: new(sync.RWMutex), //初始化sync.RWMutex
	}
	db.cond = sync.NewCond(db.mutex.RLocker())//透過db接口建立條件變數並關聯mutex的鎖
	return db //回傳內容給DataBucket
}
//build a Read method
func (db *DataBucket) Read(i int) {
	db.mutex.RLock() //透過db接口引入讀寫鎖
	defer db.mutex.RUnlock() //程序結束前釋放鎖
	var data []byte //緩衝區
	var d byte //數據
	var err error //錯誤訊息

	for {
		//d = 數據讀取結果; 若發生錯誤進入下層條件
		if d, err = db.buffer.ReadByte(); err != nil {
			if err == io.EOF{
				//若緩衝區不為nil，ptint out
				if string(data) != "" {
					fmt.Printf("read-%d: %s\n",i,data)
				}
				//若緩衝區為nil，透過Wait等待通知
				fmt.Println("read: wait")
				db.cond.Wait()
				fmt.Println("read: go")
				// data = data[:0] // 将 data 清空
				// continue
			}
		}
		//若未發生錯誤，將數據讀取結果加到緩衝區中
		fmt.Println(string(d),string(data),err)
		time.Sleep(1*time.Second)
		data = append(data, d)
		
	}
}
//build a write method 
func (db *DataBucket) Put(data []byte) (int, error)  {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	n, err := db.buffer.Write(data) //初始化變數n為buffer.Write method，並寫入一個slice緩衝區
	db.cond.Signal() //喚醒Wait狀態中的
	// db.cond.Broadcast()
	return n, err
}

func main()  {
	db := NewDataBucket()
	go db.Read(1)
	go func (i int)  {
		time.Sleep(2*time.Second)
		d := fmt.Sprintf("data-%d",i)
		db.Put([]byte(d))
	}(2)
	time.Sleep(10*time.Second)

}