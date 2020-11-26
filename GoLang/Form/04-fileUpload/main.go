package main

import (
	"os"
	"log"
	"html/template"
	"strconv"
	"io"
	"crypto/md5"
	"time"
	"fmt"
	"net/http"
)

func upload(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Method:", r.Method)
	switch r.Method {
	case "GET":
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.html")
		log.Println(t.Execute(w, token))
	case "POST":
		//解析客戶端multipart/form-data 請求 並將檔案的累計大小超出32mb的部份存入系統的臨時檔案中
		r.ParseMultipartForm(32 << 20)
		//r.FormFile取得html檔案控制代碼
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		//第一个参数为文件路径,第二个参数控制文件的打开方式,第三个参数控制文件模式
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)//儲存檔案
	}
}

func main()  {
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":9090", nil)
}