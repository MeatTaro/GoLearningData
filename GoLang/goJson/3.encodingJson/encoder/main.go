package main

import (
	"encoding/json"
	"os"
)

//將struct格式編碼為json格式，並寫入文件內

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    // 如果Child字段为nil 编码JSON时可忽略
    Child *Person `json:"child,omitempty"`
}

func main() {
	person := Person{
        Name: "John",
        Age: 40,
        Child: &Person{
            Name: "Jack",
            Age: 20,
        },
	}
	//os.Create返回的*File類型實現了io.Writer接口
	file, _ := os.Create("person.json")
	// 根据io.Writer创建Encoder 然后调用Encode()方法将对象编码成JSON
	json.NewEncoder(file).Encode(&person)
}
