package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //匯入資料庫驅動
)
//資料庫結構定義
//字首必需大寫，否則無法識別
type User struct {
	//別名為Id時，將預設為PrimeKey，並設定auto increment，如果主鍵的別名不是 id，那麼需要加上 pk 註釋，明確的說這個欄位是主鍵
	Id int
	Name string `orm:"size(100)"`
}

func init(){
	//註冊驅動
	// orm.RegisterDriver("mysql", orm.DR_MySQL)

	//設定資料庫配置資訊
	//func RegisterDataBase(aliasName, driverName, dataSource string, params ...int) error
	orm.RegisterDataBase("default","mysql","MeatTaro:9868@/test?charset=utf8",30)//30最大資料庫連線

	//RegisterModel註冊定義表單model(也可以同時註冊多個model)
	//Like orm.RegisterModel(new(User), new(Profile), new(Post))
	orm.RegisterModel(new(User))

	//同步 database 並建立表單
	//func RunSyncdb(name string, force bool, verbose bool) error
	//name means table's alias name.
	//force means run next sql if the current is error.
	//verbose means show all info when running command or not.
	orm.RunSyncdb("default", false, true)
}

func main() {
	//開啟orm資料庫連結
	//NewOrm create new orm
	o := orm.NewOrm()

	//初始化Name
	user := User{
		Name: "MeatTaro",
	}

    // Insert() 寫入表格資訊，並回傳自增id	
	//Insert(interface{}) (int64, error)
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// 重新初始化Name，Update()更新表單資料，並回傳受影響資料數量
	//Update(md interface{}, cols ...string) (int64, error)
	user.Name = "Christian"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// Read()讀取資料
	//Read(md interface{}, cols ...string) error
	u := User{
		Id: user.Id,
	}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)
    
	// 刪除表
	// delete "model" in database
	num, err = o.Delete(&user)
	fmt.Printf("NUM: %d, ERR %v\n", num, err)
}