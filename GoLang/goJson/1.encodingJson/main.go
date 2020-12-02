package main

import (
	"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
)
//讀取json文件，將內容轉為struct格式，然後更改數據
//Marshal & Unmarshal
type Person struct {
	Name	string	`json: "name"`
	Age		int	`json: "age"`
}

type Result struct {
	Group	string		`json: "group"`
	Persons	[]Person	`json: "person"`
}

func main() {
	var data Result
	//1. 讀取json文件內容，返回slice數據流
	bytes, _ := ioutil.ReadFile("data.json")
	fmt.Println("*** data.json content: ***")
	//數據流轉為string
	fmt.Println(string(bytes))
	
	
	//2. 將slice數據流Unmarshal到指定的struct格式上
	json.Unmarshal(bytes, &data)
	fmt.Println("*** unmarshal result: ***")
	fmt.Println(data)

	//3. 更改struct數據
	data.Group = "engineer"
	//將更改後的struct格式Marshal為json格式
	newBytes, _ := json.Marshal(&data)
	fmt.Println("*** update content: ***")
	fmt.Println(string(newBytes))

	//4. 使用MarshalIndent更新格式
	// newBytes, _ := json.MarshalIndent(&data, "", "       ")
	// fmt.Println("*** indent content: ***")
	// fmt.Println(string(newBytes))

	//5. 將變更結果寫入文件data.json
	ioutil.WriteFile("data.json", newBytes, os.ModeAppend)
	fmt.Println("*** write to file: data.json***")
}

