package main

import (
	"fmt"
	"os"
	"bytes"
)

//what is Buffer
//Buffer is a variable-sized buffer of bytes with Read and Write methods. The zero value for Buffer is an empty buffer ready to use.
//1. 一個字元儲存緩衝區
//2. 大小可變動
//3. 具有Read & Write 方法
//4. buffer的零值為一個待使用的緩衝區
//5. 被成功讀取數據仍留在緩衝區內

//how to use Buffer

func main()  {
	file, _ := os.Open("./test.txt")
	buf := bytes.NewBufferString("Hello World")
	buf.ReadFrom(file)
	fmt.Println(buf.String())
}


//參考資料:
//https://cloud.tencent.com/developer/article/1456243