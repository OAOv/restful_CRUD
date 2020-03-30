package repo

import (
	"encoding/json"

	"github.com/OAOv/restful_CRUD/db"
	"github.com/OAOv/restful_CRUD/types"
)

func GetUsers() ([]types.User, error) {
	var users []types.User
	result, err := db.GetUsers()
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var user types.User
		err := result.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	result.Close()
	return users, nil
}

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

func GetUser(params map[string]string) ([]types.User, error) {
	var users []types.User
	result, err := db.GetUser(params["id"])
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var user types.User
		err := result.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	result.Close()
	return users, nil
}

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
