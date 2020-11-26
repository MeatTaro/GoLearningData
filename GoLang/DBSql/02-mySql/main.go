package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"//MySQL 驅動
)

func main() {

	//sql.Open()方法用來開啟一個註冊過的資料庫驅動，這個方法回傳一個數據函式庫的Conn介面
    //第二個參數是 DSN(Data Source Name)，它是 go-sql-driver 定義的一些資料庫連結和配置資訊。(註1)
    db, err := sql.Open("mysql", "MeatTaro:9868@/test?charset=utf8")
    checkErr(err)
    defer db.Close()//退出函數時需要釋放資料庫驅動
    
    //設定資料內容
	userinfo := `CREATE TABLE IF NOT EXISTS userinfo(
        uid INT(10) NOT NULL AUTO_INCREMENT,
        username VARCHAR(64) NULL DEFAULT NULL,
        department VARCHAR(64) NULL DEFAULT NULL,
        created DATE NULL DEFAULT NULL,
        PRIMARY KEY (uid)
    );` 
    
    userdetail := `CREATE TABLE IF NOT EXISTS userdetail(
        uid INT(10) NOT NULL DEFAULT '0',
        intro TEXT NULL,
        profile TEXT NULL,
        PRIMARY KEY (uid)
    );`
	
	//建立資料表 帶入資料內容
	//Exec 函式執行 Prepare 準備好的 sql，傳入參數執行 update/insert 等操作，回傳 Result 資料
	_, err = db.Exec(userinfo)
    checkErr(err)
    
    _, err = db.Exec(userdetail)
	checkErr(err)

    //插入資料
    //Prepare() is a method，define from DB Struct
    //Prepare() create and return a prepared statement for later queries or executions.
    //Prepare() can be used to excute SQL's INSERT語句
    stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
    checkErr(err)
    //Exec()執行Prepare()準備好的SQL，傳入參數，執行update/insert，並return Resault資料
    res, err := stmt.Exec("MeatTaro", "PE", "2020-01-14")
    checkErr(err)

    //Resault 是執行 Update/Insert 等操作回傳的結果介面定義
    //LastInsertId 函式回傳由資料庫執行插入操作得到的自增 ID 號。
    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Printf("Last ID =%d\n",id)
    // 準備狀態更新userinfo表單的username的值，其值與其標定的uid皆以代數?表示，用以防止SQL注入(註2)
    stmt, err = db.Prepare("update userinfo set username=? where uid=?")
    checkErr(err)
    //輸入更新的資料，並回傳resault
    res, err = stmt.Exec("Christian", id)
    checkErr(err)
    //RowsAffected 函式回傳 query 操作影響的資料條目數。
    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Printf("Affected Num= %d\n",affect)

    //查詢資料
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Printf("id = %d\n",uid)
        fmt.Printf("username = %s\n", username)
        fmt.Printf("department = %s\n", department)
        fmt.Printf("created date = %s\n", created)
    }

    //準備狀態刪除userinfo表單的uid標定的資料列
    stmt, err = db.Prepare("delete from userinfo where uid=?")
    checkErr(err)
    //執行刪除，並回傳resault
    res, err = stmt.Exec(id)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Printf("Delete num = %d\n",affect)


}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
//註1
// user@unix(/path/to/socket)/dbname?charset=utf8
// user:password@tcp(localhost:5555)/dbname?charset=utf8
// user:password@/dbname
// user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
//註2
//https://www.itread01.com/content/1542007036.html