package main

import (
	"fmt"

	"github.com/go-xorm/xorm"
	_ "github.com/lunny/godbc"
	"time"
)

func main() {
	engine, err := xorm.NewEngine("odbc", "driver={SQL Server};SERVER=127.0.0.1;UID=sa;PWD=123;DATABASE=TodoDB")
	if err != nil {
	}
	defer engine.Close()
	engine.ShowSQL = true
	fmt.Println("init")
	//检查是否存在数据表
	// if has, _ := engine.IsTableExist(new(Task)); !has {
	// 	engine.CreateTables(new(Task))
	// }
	engine.CreateTables(new(Task))

	//插入数据
	// user := new(User)
	// user.Name = "myname"
	// affected, err := engine.Insert(user)
	fmt.Println("hello world.")

}

type Task struct {
	Id         int64
	Title      string
	Content    string
	CreateDate time.Time
	FinishDate time.Time
}

// var x *xorm.Engine

// func init() {
// 	// 创建 ORM 引擎与数据库
// 	var err error
// 	x, err = xorm.NewEngine("sqlite3", "./bank.db")
// 	if err != nil {
// 		log.Fatalf("Fail to create engine: %v\n", err)
// 	}

// 	// 同步结构体与数据表
// 	if err = x.Sync(new(Account)); err != nil {
// 		log.Fatalf("Fail to sync database: %v\n", err)
// 	}
// }
