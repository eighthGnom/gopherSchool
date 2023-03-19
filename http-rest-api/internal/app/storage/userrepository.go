package storage

import (
	"github.com/eighthGnom/http-rest-api/internal/app/models"
)

type UserRepository interface {
	Create(*models.User) error
	FindUserByEmail(string) (*models.User, error)
	FindUserByID(int) (*models.User, error)
}
