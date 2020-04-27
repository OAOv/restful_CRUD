package repository

import (
	"github.com/OAOv/restful_CRUD/types"
)

type UserRepository struct{}

func (u *UserRepository) CreateUser(user types.User) error {
	stmt, err := db.Prepare("INSERT INTO user (id, name, age) VALUES (?, ?, ?)")
	if err != nil {
		return types.ErrServerQueryError
	}
	defer stmt.Close()

	if user.ID == "" {
		user.ID = "0"
	}
	_, err = stmt.Exec(user.ID, user.Name, user.Age)
	if err != nil {
		return types.ErrServerQueryError
	}

	return nil
}

func (u *UserRepository) GetUsers() ([]types.User, error) {
	var users []types.User

	result, err := db.Query("SELECT * FROM user")
	if err != nil {
		return nil, types.ErrServerQueryError
	}
	defer result.Close()

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

func (u *UserRepository) GetUser(id string) (types.User, error) {
	var user types.User
	result, err := db.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		return user, types.ErrServerQueryError
	}
	defer result.Close()

	result.Next()
	err = result.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return user, types.ErrNotFound
	}

	return user, nil
}

func (u *UserRepository) UpdateUser(id string, user map[string]interface{}) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return types.ErrServerQueryError
	}
	result.Next()
	err = result.Scan()
	if err != nil {
		return types.ErrNotFound
	}
	result.Close()

	isFirst := true
	sql := "UPDATE user SET"
	for key, value := range user {
		if isFirst {
			sql += " " + key + " = \"" + value.(string) + "\""
			isFirst = false
		} else {
			sql += ", " + key + " = \"" + value.(string) + "\""
		}
	}
	sql += " WHERE id = " + id
	_, err = tx.Exec(sql)
	if err != nil {
		tx.Rollback()
		return types.ErrServerQueryError
	}

	sql = "UPDATE record SET user_name = (SELECT name FROM user WHERE id = \"" + id + "\") WHERE user_id = \"" + id + "\""
	_, err = tx.Exec(sql)
	if err != nil {
		tx.Rollback()
		return types.ErrServerQueryError
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

//transaction
func (u *UserRepository) DeleteUser(id string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM record WHERE user_id = ?", id)
	if err != nil {
		tx.Rollback()
		return types.ErrServerQueryError
	}

	_, err = tx.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return types.ErrServerQueryError
	}

	err = tx.Commit()
	if err != nil {
		return nil
	}

	return nil
}
