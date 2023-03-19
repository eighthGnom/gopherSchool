package postgresstorage_test

import (
	"testing"

	"github.com/eighthGnom/http-rest-api/internal/app/models"
	"github.com/eighthGnom/http-rest-api/internal/app/storage"
	"github.com/eighthGnom/http-rest-api/internal/app/storage/postgresstorage"
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

func TestUserRepository_FindUserByID(t *testing.T) {
	store, teardown := postgresstorage.TestStorage(t, databaseURL)
	defer teardown(userTable)
	testUser := models.TestUser(t)
	err := store.User().Create(testUser)
	if err != nil {
		t.Fatal(err)
	}

	user, err := store.User().FindUserByID(999)
	assert.Empty(t, user)
	assert.EqualError(t, err, storage.ErrRecordNotFound.Error())

	user, err = store.User().FindUserByID(testUser.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}
