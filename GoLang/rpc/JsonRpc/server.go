package main

import (
	"net/rpc/jsonrpc"
	"net"
	"net/rpc"
	"log"
	"JsonRpc/utils"
)

type ServiceHeadler struct {}

//GetName方法透過傳入Item指針，改變Item結構裡的值
func (t *ServiceHeadler) GetName(id int, item *utils.Item) error  {
	log.Printf("receive GetName call, id: %d", id)
	item.Id = id
	item.Name = "MeatTaro"
	return nil
}
//SaveName透過
func (t *ServiceHeadler) SaveName(item utils.Item, resp *utils.Response) error {
	log.Printf("receive SaveName call, Item: %v", item)
	resp.Ok = true
	resp.Id = item.Id
	resp.Message = "Save Success"
	return nil
}

func main()  {
	//NewServer()初始化一個新的Server
	server := rpc.NewServer()
	//Listen announces in the local network address
	//return error and a Listener type
	//回傳的Listener是網路通用的協議監聽介面
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("started listen server fail: %v", err)
	}
	defer listener.Close()

	log.Println("Start listen on port localhost: 8080")
	//初始化服務端struct
	serviceHeadler := &ServiceHeadler{}
	//透過rpc服務註冊type ServiceHeadler struct {}
	err = server.Register(serviceHeadler)
	if err != nil {
		log.Fatal("Registed server fail: %v", err)
	}

	for {
		//Listener介面底下的Accept方法等待接收並回傳連結給Listener
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("accept client fail: %v", err)
		}
		//ServeCodec為隸屬於type Server的method，可使用指定的編(解)碼器進行request的解碼與response編碼
		//func (server *Server) ServeCodec(codec ServerCodec)
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
//參考資料:
//https://golang.google.cn/pkg/net/#Listener
//https://golang.google.cn/pkg/net/rpc/#Server