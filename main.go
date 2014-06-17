package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	//"strings"

	"./DAL"
	"time"
)

func main() {
	rows := DAL.GetRows("SELECT * FROM " + "tmpTable")

	var id int
	var name string
	for rows.Next() {
		if err := rows.Scan(&id, &name); err == nil {
			fmt.Println(id, " ", name)
		}

	}
	// rows := QuerryData("tmpTable")

	// var id string
	// var name string
	// for rows.Next() {
	// 	if err := rows.Scan(&id, &name); err == nil {
	// 		fmt.Println("id:", id, " name:", name)
	// 	}
	// }
	//fmt.Println("什么垃圾go啊, 配置个运行环境都死人了")
}

func QuerryData(table string) *sql.Rows {
	db, err := sql.Open("sqlite3", "./odb.s3db")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM " + table)
	checkErr(err)

	db.Close()

	return rows
}

func QueryPageData(index int) string {

	rows := QuerryData("pageinfo")

	var id int
	var data string
	page := ""

	for rows.Next() {
		if err := rows.Scan(&id, &data); err == nil && id == index {
			page = data
		}
	}

	return page
}

func QueryDateData() map[int]string {

	rows := QuerryData("dateinfo")

	var date int
	var data string

	memoryCache := make(map[int]string)
	for rows.Next() {
		if err := rows.Scan(&date, &data); err == nil {
			memoryCache[date] = data
		}
	}

	return memoryCache
}

type Task struct {
	Id         int64
	Title      string
	Content    string
	CreateDate time.Time
	FinishDate time.Time
}

func checkErr(err error) {
	if err != nil {
		//panic(err)
		fmt.Println("panic: " + err.Error())
	}
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
