package repo

import (
	"context"
	"encoding/json"

	"github.com/OAOv/restful_CRUD/db"
)

func getUsers(ctx context.Context) {
	var users []User

	result, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var user User
		err := result.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}
