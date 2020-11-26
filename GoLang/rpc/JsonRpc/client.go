package main

import (
	"JsonRpc/utils"
	"net/rpc/jsonrpc"
	"log"
	"time"
	"net"
)

func main()  {
	conn, err := net.DialTimeout("tcp", "localhost:8080", 30 *time.Second)
	if err != nil {
		log.Fatalf("客户端发起连接失败：%v", err)
	}
	defer conn.Close()

	client := jsonrpc.NewClient(conn)
	var item utils.Item
	client.Call("ServiceHeadler.GetName", 1, &item)
	log.Printf("ServiceHeadler.GetName 返回結果: %v\n", item)

	var resp utils.Response
	item = utils.Item{2, "MeatTaro2"}
	client.Call("ServiceHeadler.SaveName", item, &resp)
	log.Printf("ServiceHeadler.SaveName 返回結果: %v\n", resp)
}