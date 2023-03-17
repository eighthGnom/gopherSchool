package postgresstorage

import (
	"database/sql"

	"github.com/eighthGnom/http-rest-api/storage"
	_ "github.com/lib/pq"
)

type Storage struct {
	db             *sql.DB
	userRepository storage.UserRepository
}

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) User() storage.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		storage: s,
	}
	return s.userRepository
}
