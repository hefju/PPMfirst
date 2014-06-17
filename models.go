package main

//import (
//	// "errors"
//	// "log"

//	"github.com/go-xorm/xorm"
//	_ "github.com/mattn/go-sqlite3"
//)

//type Task struct {
//	Id         int64
//	Title      string
//	Content    string
//	CreateDate time.Time
//	FinishDate time.Time
//}

//var x *xorm.Engine

//func init() {
//	// 创建 ORM 引擎与数据库
//	var err error
//	x, err = xorm.NewEngine("sqlite3", "./bank.db")
//	if err != nil {
//		log.Fatalf("Fail to create engine: %v\n", err)
//	}

//	// 同步结构体与数据表
//	if err = x.Sync(new(Account)); err != nil {
//		log.Fatalf("Fail to sync database: %v\n", err)
//	}
//}
