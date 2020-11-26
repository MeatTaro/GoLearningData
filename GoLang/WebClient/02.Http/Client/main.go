package main

import (
	"os"
	"io"
	"fmt"
	"net/http"
)

func main()  {
	//http.Request底下的NewRequest方法初始化Requester
	//func NewRequest(method, url string, body io.Reader) (*Request, error)
	req, err := http.NewRequest("Get","https://xueyuanjun.com",nil)

	if err != nil {
		fmt.Printf("請求初始化失敗:%v", err)
		return
	}
	//使用Header底下的Add方法新增Head鍵值對
	//func (h Header) Add(key, value string)
	req.Header.Add("Custom-Header", "Custom-Value")
	//設置客戶端接口
	client := &http.Client {}

	//啟動客戶端伺服器
	//func (c *Client) Do(req *Request) (*Response, error)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("客户端发起请求失败：%v", err)
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}