package DAL

import (
	"database/sql"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	driverName     = "sqlite3"
	dataSourceName = "./odb.s3db"
)

//----------------------------xorm-------------------
type User struct {
	Id   int64  `xorm:autoincr`
	Name string `xorm:"varchar(25) not null unique 'usr_name'"`
}

func CreateTableByXorm() {
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	checkErr(err)
	defer engine.Close()
	engine.ShowSQL = true

	engine.Sync(new(User))
	// iamvoid(affected)

}
func getEngin() *xorm.Engine {
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	checkErr(err)
	return engine
}

func DeleteByXorm() {
	sql := "delete from user"
	_, err := getEngin().Exec(sql)
	checkErr(err)
}

func InsertByXorm(name string) int64 {

	user := new(User)
	user.Name = name
	affected, err := getEngin().Insert(user)
	checkErr(err)
	return affected
}

func FindByIDByXorm(uid int) []User {
	users := make([]User, 0)
	err := getEngin().Where("id=?", uid).Find(&users)
	checkErr(err)
	return users
}

func FindAllByXorm() []User {
	users := make([]User, 0)
	err := getEngin().Where("1=1").Find(&users)
	checkErr(err)
	return users
}

func iamvoid(x interface{}) {

}

//---------------------------------database/sql--------------

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
