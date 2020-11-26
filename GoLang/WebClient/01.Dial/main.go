package main
//Package net provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets.
//Although the package provides access to low-level networking primitives, most clients will need only the basic interface provided by the Dial, Listen, and Accept functions and the associated Conn and Listener interfaces. The crypto/tls package uses the same interfaces and similar Dial and Listen functions.

//Dial(): Dial connects to the address on the named network.(server)
// func Dial(network, address string) (Conn, error) {
//     var d Dialer
//     return d.Dial(network, address)
// }

//type Coon: Conn is a generic stream-oriented network connection.
//Multiple goroutines may invoke methods on a Conn simultaneously.

//network: 網路協議(TCP、UDP、ICMP)
//address: IP位址

//1.TCP connect: coon, err := net.Dial("tcp","192.168.10.10.80")
//2.UDP connect: conn, err := net.Dial("udp", "192.168.10.10:8888")
//3.ICMP connect(name): conn, err := net.Dial("ip4:icmp", "www.xueyuanjun.com")
//4.ICMP connect(numbers) : conn, err := net.Dial("ip4:1", "10.0.0.3")


import (
	"io"
	"bytes"
	"net"
	"fmt"
	"os"
)


func main()  {
	if len(os.Args) != 2 { //參數列[1]不可為空
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	//從參數列[1]讀取訊息
	service := os.Args[1]

	//建立網路連接，回傳連接對象(coon)
	coon, err := net.Dial("tcp", service)

	//連接錯誤則傳入err執行錯誤func checkError()
	checkError(err)

	//調用連接對象(coon)提供的 method Write 發送request
	_, err = coon.Write([]byte ("HEAD / HTTP/1.0\r\n\r\n"))

	//通過連接對象(coon)提供的 method Read 讀取response
	result, err := readFully(coon)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error)  {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(coon net.Conn) ([]byte, error) {
	//讀取所有response後關閉連結
	defer coon.Close()
	//建立緩衝區
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	//回傳coon直到所有內容傳送完畢
	for {
		//n loading all of buf
		n, err := coon.Read(buf[0:])
		//將buf所有內容寫入result
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF{
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

// 參考資料:
//https://xueyuanjun.com/post/20970
//https://golang.org/pkg/net/#Conn
//https://gobyexample.com/command-line-arguments
//https://openhome.cc/Gossip/Go/StdOutInErr.html