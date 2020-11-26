package main

import (
	"fmt"
	"encoding/json"
)

func main()  {
	u := []byte(`{"name": "学院君", "website": "https://xueyuanjun.com", "age": 18, "male": true, "skills": ["Golang", "PHP"]}`)
	var usr interface{}
	err := json.Unmarshal(u, &usr)
	if err != nil {
		fmt.Printf("fail: %v \n", err)
		return
	}
	fmt.Printf("%#v\n", usr)

	//訪問解碼後數據
	usrx, ok := usr.(map[string]interface{})
	fmt.Println(ok)
	if ok {
		for k, v :=range usrx {
			switch vx := v.(type) {
			case string:
				fmt.Println(k, vx)
			case int:
				fmt.Println(k, vx)
			case bool:
				fmt.Println(k, vx)
			case []interface{}:
				fmt.Println(k)
				for i, j := range vx {
					fmt.Println(i, j)
				}
			default:
				fmt.Println(k, "is another type not handle yet")
			}
		}
	}
}