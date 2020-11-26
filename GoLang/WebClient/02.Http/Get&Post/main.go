package main
//golang的標準api中用於http客戶端請求的主要有三個api :
//http.Get	發送get請求
//http.Post	post請求提交指定類型的數據
//http.PostForm	post請求提交application/x-www-form-urlencoded數據
import (
	// "strings"
	"os"
	"io"
	"fmt"
	"net/http"
)

func main()  {
	//通過Get(URL)發起請求，回傳 responce 與 err
	resp, err := http.Get("https://xueyuanjun.com")

	//Post方法為新增資料，分別傳入(URL(含API),資料格式,資料內容)，以下範例
	//URL(API):要求上傳用戶頭像
	//資料格式:jpeg
	//資料內容:圖像位址
	//resp, err := http.Post("https://xueyuanjun.com/avatar", "image/jpeg", &imageDataBuf)

	//PostForm與Post差異為透過 url.Values 進行封裝表單參數
	//resp, err := http.PostForm("https://xueyuanjun.com/login", url.Values{"name":{"学院君"}, "password": {"test-passwd"}}) 
		
	if err != nil {
		fmt.Printf("fail: %v", err)
		return
	}
	//於程式最後釋放資源
	defer resp.Body.Close()
	// 透過resp.Body獲取body
	file, err := os.Create("./proverbs.txt")
	defer file.Close()
	io.Copy(file, resp.Body)

	
}

//參考資料
//https://learnku.com/articles/23430/golang-learning-notes-1-http-client-foundation
//https://golang.org/pkg/net/http/#Post