package teststorage_test

import (
	"testing"

	"github.com/eighthGnom/http-rest-api/internal/app/models"
	"github.com/eighthGnom/http-rest-api/internal/app/storage"
	"github.com/eighthGnom/http-rest-api/internal/app/storage/teststorage"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	store := teststorage.New()
	testUser := models.TestUser(t)
	err := store.User().Create(testUser)

	assert.NotEmpty(t, testUser.ID)
	assert.NotEmpty(t, testUser.EnscriptedPassword)
	assert.NoError(t, err)
}

func TestUserRepository_FindUserByEmail(t *testing.T) {
	store := teststorage.New()
	testUser := models.TestUser(t)
	user, err := store.User().FindUserByEmail(testUser.Email)
	assert.Nil(t, user)
	assert.EqualError(t, err, storage.ErrRecordNotFound.Error())

	err = store.User().Create(testUser)
	if err != nil {
		t.Fatal(err)
	}
	user, err = store.User().FindUserByEmail(testUser.Email)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestUserRepository_FindUserByID(t *testing.T) {
	testStorage := teststorage.New()
	testUser := models.TestUser(t)
	err := testStorage.User().Create(testUser)
	if err != nil {
		t.Fatal(err)
	}
	user, err := testStorage.User().FindUserByID(999)
	assert.Empty(t, user)
	assert.EqualError(t, err, storage.ErrRecordNotFound.Error())

	user, err = testStorage.User().FindUserByID(testUser.ID)
	assert.NotEmpty(t, user)
	assert.NoError(t, err)
}
