package storage_test

import (
	"testing"

	"github.com/eighthGnom/http-rest-api/models"
	"github.com/eighthGnom/http-rest-api/storage"
	"github.com/stretchr/testify/assert"
)

var userTable = "users"

func TestUserRepository_Create(t *testing.T) {
	s, teardown := storage.TestStorage(t, databaseURL)
	defer teardown(userTable)
	user, err := s.User().CreateUser(&models.User{
		Email: "test@gmail.com",
	})
	assert.NotNil(t, user)
	assert.NoError(t, err)

}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := storage.TestStorage(t, databaseURL)
	defer teardown(userTable)
	email := "test@gmail.com"
	u1, err := s.User().FindUserByEmail(email)
	assert.Nil(t, u1)
	assert.Error(t, err)

	_, err = s.User().CreateUser(&models.User{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	u2, err := s.User().FindUserByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
