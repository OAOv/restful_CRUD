package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func OpenDB() *sql.DB {
	db, err = sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	return db
}

func GetUsers() *sql.Rows {
	result, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}

	return result
}

func CreateUser(name string, age string) *sql.Stmt {
	stmt, err := db.Prepare("INSERT INTO user (name, age) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(name, age)
	if err != nil {
		panic(err.Error())
	}

	return stmt
}

func GetUser(id string) *sql.Rows {
	result, err := db.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	return result
}

func UpdateUser(id string, name string, existsName bool, age string, existsAge bool) *sql.Stmt {
	stmt, err := db.Prepare("UPDATE user SET name  = ?, age = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	if !existsName {
		stmt, err = db.Prepare("UPDATE user SET age = ? WHERE id = ?")
		_, err = stmt.Exec(age, id)
		if err != nil {
			panic(err.Error())
		}
	} else if !existsAge {
		stmt, err = db.Prepare("UPDATE user SET name = ? WHERE id = ?")

		_, err = stmt.Exec(name, id)
		if err != nil {
			panic(err.Error())
		}
	} else {
		_, err = stmt.Exec(name, age, id)
		if err != nil {
			panic(err.Error())
		}
	}

	return stmt
}

func DeleteUser(id string) *sql.Stmt {
	stmt, err := db.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	return stmt
}
