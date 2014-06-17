package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	//"strings"

	"./DAL"
	"time"
)

func displayAll() {
	list := DAL.FindAllByXorm()
	for i, v := range list {
		fmt.Println("row:", i, " ", v.Id, "_", v.Name)
	}
}
func main() {

	//DAL.CreateTableByXorm()
	//DAL.InsertByXorm("csc")
	//DAL.DeleteByXorm()
	displayAll()

	// list := DAL.FindByIDByXorm(2)
	// for i, v := range list {
	// 	fmt.Println("row:", i, " ", v.Id, "_", v.Name)
	// }

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
