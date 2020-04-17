package repository

import (
	"database/sql"

	"github.com/OAOv/restful_CRUD/types"
)

type UserRepository struct{}

var db *sql.DB
var err error

func OpenDB() (*sql.DB, error) {
	db, err = sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/test")
	return db, err
}

func CreateUser(user types.User) error {
	stmt, err := db.Prepare("INSERT INTO user (id, name, age) VALUES (?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		return types.ErrServerQueryError
	}
	if user.ID == "" {
		user.ID = "0"
	}
	_, err = stmt.Exec(user.ID, user.Name, user.Age)
	if err != nil {
		return types.ErrServerQueryError
	}

	return nil
}

func GetUsers() ([]types.User, error) {
	var users []types.User

	result, err := db.Query("SELECT * FROM user")
	defer result.Close()
	if err != nil {
		return nil, types.ErrServerQueryError
	}

	for result.Next() {
		var user types.User
		err := result.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, types.ErrInvalidParams
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUser(id string) (types.User, error) {
	var user types.User
	result, err := db.Query("SELECT * FROM user WHERE id = ?", id)
	defer result.Close()
	if err != nil {
		return user, types.ErrServerQueryError
	}

	result.Next()
	err = result.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return user, types.ErrNotFound
	}

	return user, nil
}

func UpdateUser(user types.User) error {
	stmt, err := db.Prepare("UPDATE user SET name  = ?, age = ? WHERE id = ?")
	defer stmt.Close()

	if user.Name != "" && user.Age == "" {
		stmt, err = db.Prepare("UPDATE user SET name  = ? WHERE id = ?")
		_, err = stmt.Exec(user.Name, user.ID)
	} else if user.Name == "" && user.Age != "" {
		stmt, err = db.Prepare("UPDATE user SET age = ? WHERE id = ?")
		_, err = stmt.Exec(user.Age, user.ID)
	} else {
		_, err = stmt.Exec(user.Name, user.Age, user.ID)
	}

	if err != nil {
		return types.ErrServerQueryError
	}

	return nil
}

func DeleteUser(id string) error {
	stmt, err := db.Prepare("DELETE FROM user WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		return types.ErrServerQueryError
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return types.ErrServerQueryError
	}

	return nil
}
