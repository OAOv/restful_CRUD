package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func OpenDB() (*sql.DB, error) {
	db, err = sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/test")
	//db, err = sql.Open("mysql", "root:0000@tcp(192.168.99.100:3306)/test") //win10 home docker toolbox, docker-machine ip
	return db, err
}

func GetUsers() (*sql.Rows, error) {
	result, err := db.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}

	return result, err
}

func CreateUser(name string, age string) (*sql.Stmt, error) {
	stmt, err := db.Prepare("INSERT INTO user (name, age) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(name, age)
	if err != nil {
		return nil, err
	}

	return stmt, err
}

func GetUser(id string) (*sql.Rows, error) {
	result, err := db.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return result, err
}

func UpdateUser(id string, name string, existsName bool, age string, existsAge bool) (*sql.Stmt, error) {
	stmt, err := db.Prepare("UPDATE user SET name  = ?, age = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	if !existsName {
		stmt, err = db.Prepare("UPDATE user SET age = ? WHERE id = ?")
		_, err = stmt.Exec(age, id)
		if err != nil {
			return nil, err
		}
	} else if !existsAge {
		stmt, err = db.Prepare("UPDATE user SET name = ? WHERE id = ?")

		_, err = stmt.Exec(name, id)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = stmt.Exec(name, age, id)
		if err != nil {
			return nil, err
		}
	}

	return stmt, err
}

func DeleteUser(id string) (*sql.Stmt, error) {
	stmt, err := db.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return nil, err
	}

	return stmt, err
}
