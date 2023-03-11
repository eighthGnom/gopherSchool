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
	testUser := models.TestUser(t)
	user, err := s.User().CreateUser(testUser)
	assert.NotNil(t, user)
	assert.NoError(t, err)

}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := storage.TestStorage(t, databaseURL)
	defer teardown(userTable)
	testUser := models.TestUser(t)
	u1, err := s.User().FindUserByEmail(testUser.Email)
	assert.Nil(t, u1)
	assert.Error(t, err)

	_, err = s.User().CreateUser(testUser)
	if err != nil {
		t.Fatal(err)
	}
	u2, err := s.User().FindUserByEmail(testUser.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
