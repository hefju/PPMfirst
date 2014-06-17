package DAL

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const (
	driverName     = "sqlite3"
	dataSourceName = "./odb.s3db"
)

func getDB() *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	checkErr(err)
	return db
}
func GetRows(sql string) *sql.Rows {

	rows, err := getDB().Query(sql)
	checkErr(err)
	return rows
}

// func ExecSQL(sql string) {
// 	db = getDB()

// }

func checkErr(err error) {
	if err != nil {
		//panic(err)
		fmt.Println("panic: " + err.Error())
	}
}
