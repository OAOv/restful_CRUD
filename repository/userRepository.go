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

/*
func CreateUser(body []byte) error {
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	name := keyVal["name"]
	age := keyVal["age"]

	stmt, err := db.CreateUser(name, age)
	if err != nil {
		return err
	}

	defer stmt.Close()
	return nil
}
*/
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
			return nil, types.ErrInvalidType
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
		return user, types.ErrInvalidType
	}

	return user, nil
}

/*
func UpdateUser(params map[string]string, body []byte) error {
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newName, existsName := keyVal["name"]
	newAge, existsAge := keyVal["age"]

	stmt, err := db.UpdateUser(params["id"], newName, existsName, newAge, existsAge)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func DeleteUser(parmas map[string]string) error {
	stmt, err := db.DeleteUser(parmas["id"])
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
*/
