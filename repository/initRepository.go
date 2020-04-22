package repository

import "database/sql"

var db *sql.DB
var err error

func OpenDB() (*sql.DB, error) {
	db, err = sql.Open("mysql", "root:0000@tcp(192.168.99.100:3306)/test")
	return db, err
}
