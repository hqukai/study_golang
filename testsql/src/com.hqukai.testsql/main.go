package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	//"time"
)

func main() {
	db, err := sql.Open("mysql", " root:root@tcp(192.168.77.107:3306)/test?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("hankai", "研发部", "2013-6-23")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username = ? where uid = ?")
	checkErr(err)

	res, err = db.Exec("hankai1", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
