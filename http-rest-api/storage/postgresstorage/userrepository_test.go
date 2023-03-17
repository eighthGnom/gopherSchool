package postgresstorage_test

import (
	"testing"

	"github.com/eighthGnom/http-rest-api/models"
	"github.com/eighthGnom/http-rest-api/storage"
	"github.com/eighthGnom/http-rest-api/storage/postgresstorage"
	"github.com/stretchr/testify/assert"
)

var userTable = "users"

func TestUserRepository_Create(t *testing.T) {
	s, teardown := postgresstorage.TestStorage(t, databaseURL)
	defer teardown(userTable)
	testUser := models.TestUser(t)
	err := s.User().Create(testUser)
	assert.NotEmpty(t, testUser.ID)
	assert.NotEmpty(t, testUser.EnscriptedPassword)
	assert.NoError(t, err)

}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := postgresstorage.TestStorage(t, databaseURL)
	defer teardown(userTable)
	testUser := models.TestUser(t)
	user, err := s.User().FindUserByEmail(testUser.Email)
	assert.Nil(t, user)
	assert.EqualError(t, err, storage.ErrRecordNotFound.Error())

	err = s.User().Create(testUser)
	if err != nil {
		t.Fatal(err)
	}
	user, err = s.User().FindUserByEmail(testUser.Email)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
