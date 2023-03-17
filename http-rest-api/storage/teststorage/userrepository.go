package teststorage

import (
	"github.com/eighthGnom/http-rest-api/models"
	"github.com/eighthGnom/http-rest-api/storage"
)

type UserRepository struct {
	storage *Storage
	db      map[string]*models.User
}

func (ur *UserRepository) Create(user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	if err := user.EnscriptPassword(); err != nil {
		return err
	}
	user.ID = len(ur.db) + 1
	ur.db[user.Email] = user

	return nil
}

func (ur *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	user, ok := ur.db[email]
	if !ok {
		return nil, storage.ErrRecordNotFound
	}
	return user, nil

}
