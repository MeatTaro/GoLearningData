package main

//net/rpc 包允许 RPC 客户端程序通过网络或是其他 I/O 连接调用一个服务端对象的公开方法（大小字母开头）。在 RPC 服务端，需要将这个对象注册为可访问的服务，之后该对象的公开方法就能够以远程的方式提供访问。

import (
	"net/http"
	"log"
	"net"
	"net/rpc"
	"errors"
	"rpc/utils"
)

type MathService struct {

}

func (m *MathService) Multiply(args *utils.Args, reply *int) error  {
	*reply = args.A * args.B
	return nil
}

func (m *MathService) Divide(args *utils.Args, reply *int) error {
	if args.B == 0 {
		return errors.New("除数不能为0")
	}
	*reply = args.A / args.B
	return nil
}

func main() {
	math := new(MathService)

	//Register 註冊服務端裡接收request的方法
	rpc.Register(math)
	//HandleHTTP宣告服務端裡接收處理request的方法
	rpc.HandleHTTP()
	//宣告network address，用於建立客戶端接口
	//Listen announces in the local network address
	//return error and a Listener type
	//回傳的Listener是網路通用的協議監聽器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("启动服务监听失败:", err)
	}
	//啟動HTTP伺服器
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("启动 HTTP 服务失败:", err)
	}
}

//參考資料:
//https://xueyuanjun.com/post/21069
//https://colobu.com/2016/09/18/go-net-rpc-guide/
//https://golang.google.cn/pkg/net/rpc/#Register

//**https://doc.rpcx.io/