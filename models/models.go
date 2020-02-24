package models

import (
	"database/sql"
)

func Sqlconn()(*sql.DB){
	db, err := sql.Open("mysql", "root:root@/simplecomment?charset=utf8")
	if err != nil {
		panic(err)
		return nil
	}
	return db
}
func SqlExe(sql string) bool {

	db:=Sqlconn()
	defer db.Close()
	println(sql)
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
		return false

	}
	res, err := stmt.Exec()
	if err != nil {
		panic(err)
		return false

	}
	line,_:=res.RowsAffected()
	println(line)
	if line==1{
		return true
	}else{
		return false
	}
}
func SqlQuery(sql string) *sql.Rows{
	db:=Sqlconn()
	defer db.Close()
	println(sql)

	rows, err := db.Query("SELECT * FROM comment")
	if err != nil {
		panic(err)
		return nil

	}
	return rows
}