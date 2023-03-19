package teststorage

import (
	"github.com/eighthGnom/http-rest-api/internal/app/models"
	"github.com/eighthGnom/http-rest-api/internal/app/storage"
)

type UserRepository struct {
	storage *Storage
	db      map[int]*models.User
}

func (ur *UserRepository) Create(user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	if err := user.EnscriptPassword(); err != nil {
		return err
	}
	user.ID = len(ur.db) + 1

	ur.db[user.ID] = user

	return nil
}

func (ur *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	for _, user := range ur.db {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, storage.ErrRecordNotFound

}

func (ur *UserRepository) FindUserByID(id int) (*models.User, error) {
	user, ok := ur.db[id]
	if !ok {
		return nil, storage.ErrRecordNotFound
	}
	return user, nil
}
