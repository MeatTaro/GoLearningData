package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)
//***於func內轉換格式***
// var JSON = `{
// 	"group": "programmer",
// 	"persons": [
// 	  {
// 		"name": "Jack",
// 		"age": 25
// 	  },
// 	  {
// 		"name": "Lily",
// 		"age": 20
// 	  }
// 	]
//   }`

// 使用map格式 将JSON字符串转为对象

func main() {
	var data map[string]interface{}
	bytes, _ := ioutil.ReadFile("data.json")
	// var bytes []byte
	// bytes = []byte(JSON)

	//將slice數據流Unmarshal到map上
	//json格式的 title:value 對應 map為 string:interface{}
	json.Unmarshal(bytes, &data)
	fmt.Println("group:", data["group"])

	//轉化為[]interface{}類型
	person := data["persons"].([]interface{})

	for index, item := range person {
		// 類型轉換
		person := item.(map[string]interface{})
		age := person["age"]
		// 更改年龄
        person["age"] = age.(float64) + 1

        // 打印最新个人信息
        fmt.Println("person", index, ":", person["name"].(string), age, "->", person["age"].(float64))
    }

    // 序列化为JSON字符串
    newBytes, _ := json.MarshalIndent(&data, "", "  ");

    // 打印新的JSON数据
    fmt.Println(string(newBytes))
}