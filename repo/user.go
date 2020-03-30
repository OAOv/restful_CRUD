package repo

import (
	"encoding/json"

	"github.com/OAOv/restful_CRUD/db"
	"github.com/OAOv/restful_CRUD/types"
)

func GetUsers() []types.User {
	var users []types.User
	result := db.GetUsers()

	for result.Next() {
		var user types.User
		err := result.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	result.Close()
	return users
}

func CreateUser(body []byte) {
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	name := keyVal["name"]
	age := keyVal["age"]

	stmt := db.CreateUser(name, age)

	defer stmt.Close()
	return
}

func GetUser(params map[string]string) []types.User {
	var users []types.User
	result := db.GetUser(params["id"])

	for result.Next() {
		var user types.User
		err := result.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	result.Close()
	return users
}

func UpdateUser(params map[string]string, body []byte) {
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newName, existsName := keyVal["name"]
	newAge, existsAge := keyVal["age"]

	stmt := db.UpdateUser(params["id"], newName, existsName, newAge, existsAge)
	defer stmt.Close()
	return
}

func DeleteUser(parmas map[string]string) {
	stmt := db.DeleteUser(parmas["id"])
	defer stmt.Close()
	return
}
