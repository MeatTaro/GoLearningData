package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	db, err := sql.Open("mysql", "MeatTaro:9868@/test?charset=utf8")
	checkErr(err)
	defer db.Close()

	//設定資料內容
	tableUser := `CREATE TABLE IF NOT EXISTS users(
    	username VARCHAR(64) NULL DEFAULT NULL,
		password VARCHAR(64) NULL DEFAULT NULL
	);` 
	
	//建立資料表 帶入資料內容
	//Exec 函式執行 Prepare 準備好的 sql，傳入參數執行 update/insert 等操作，回傳 Result 資料
	_, err = db.Exec(tableUser)
	checkErr(err)
	
	//資料表欄位插入數據
	rs, err := db.Exec(	"INSERT INTO users(username, password) VALUES ('MeatTaro','9868')")
	checkErr(err)

	//通過RowsAffected獲取受影響的行數
	rowCount, err := rs.RowsAffected()
	checkErr(err)
	log.Printf("inserted %d rows", rowCount)
	
	//使用Query方法執行sql的SELECT語句（查詢）返回一個Rows(Result of Query)
	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)
	
	//Next()依序準備 result Row 供 Scan() 寫入
	//Next() 必須區分 false 是由於所有的result Row已Scan()完畢 還是發生錯誤而執行錯誤訊息
	for rows.Next() {
		var username string
		var password string
		//Scan() copies result Row裡的正確欄位的value of point 給對應的變數
		err = rows.Scan(&username, &password)
		checkErr(err)
		fmt.Println(username, password)
	}
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}