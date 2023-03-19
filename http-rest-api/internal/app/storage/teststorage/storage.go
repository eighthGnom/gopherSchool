package teststorage

import (
	"github.com/eighthGnom/http-rest-api/internal/app/models"
	"github.com/eighthGnom/http-rest-api/internal/app/storage"
)

type Storage struct {
	userRepository storage.UserRepository
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) User() storage.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		storage: s,
		db:      make(map[int]*models.User),
	}
	return s.userRepository
}
