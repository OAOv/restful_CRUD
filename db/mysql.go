package db

import "database/sql"

func OpenDB() {
	db, err := sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
