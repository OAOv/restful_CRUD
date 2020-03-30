package db

import "database/sql"

type MySQL struct {
	db  *sql.DB
	err error
}

var mysql *MySQL

func OpenDB() *MySQL {
	mysql.db, mysql.err = sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/test")

	if mysql.err != nil {
		panic(mysql.err.Error())
	}
	defer mysql.db.Close()

	return mysql
}

func getUsers() *sql.Rows {
	result, err := mysql.db.Query("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()
	return result
}
