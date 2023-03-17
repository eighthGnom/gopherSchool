package storage

import "github.com/eighthGnom/http-rest-api/models"

type UserRepository interface {
	Create(*models.User) error
	FindUserByEmail(string) (*models.User, error)
}
