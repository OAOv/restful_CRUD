package repo

import (
	"context"

	"github.com/OAOv/restful_CRUD/db"
	"github.com/OAOv/restful_CRUD/types"
)

type MySQLRepo struct {
	repo *db.MySQL
}

func NewRepo(mysql *db.MySQL) (repo *MySQLRepo) {
	return &MySQLRepo{
		repo: mysql,
	}
}

func (m *MySQLRepo) getUsers(ctx context.Context) []types.User {
	var users []types.User

	result, err := m.repo.db.Query("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var user types.User
		err := result.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	return users
}
