package storage

import (
	"fmt"

	"github.com/eighthGnom/http-rest-api/models"
)

var userTable = "users"

type UserRepository struct {
	storage *Storage
}

func (ur *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	err = user.EnscriptPassword()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("insert into %s (email, enscripted_password) values ($1, $2) returning id", userTable)
	err = ur.storage.db.QueryRow(query, user.Email, user.EnscriptedPassword).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	query := fmt.Sprintf("select id, email, enscripted_password from %s where email = $1", userTable)
	err := ur.storage.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.EnscriptedPassword)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
