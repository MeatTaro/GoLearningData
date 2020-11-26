package main

import (
	"errors"
	"fmt"
)

type errUserNameExist struct{
	UserName string
}

func (e errUserNameExist) Error() string{
	return fmt.Sprintf("username %s allready exist",e.UserName)
}

func isErrUserNameExist(err error) bool{
	_, ok := err.(errUserNameExist)
	return ok
}
// 確認使用者名稱是否存在
func checkUserNameExist(username string) (bool, error){
//函式名稱 checkUserNameExist 
//定義傳入值(名稱username 型態為字串string) 
//定義多重回傳值型態(布林型, 錯誤型)
	if username == "MeatTaro"{
		return true, errUserNameExist{UserName:username}
	}
 
	return false, nil
	//如果使用者名稱為 "MeatTaro"
	//回傳 true, 以fmt.Errorf印出錯誤字串	
	//否則回傳 false, error為空值
}

func main()  {
	//宣告變數err為func並傳入"MeatTaro"後傳回2個值(bool,error)
//我們並未使用到bool，以下底線省略
	if _, err := checkUserNameExist("PualLiLi"); err != nil{
		fmt.Println(err)
	} 
}