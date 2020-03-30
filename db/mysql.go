package db

import "database/sql"

func OpenDB(pass string) {
	db, err := sql.Open("mysql", "root:"+pass+"@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
