package models_test

import (
	"testing"

	"github.com/eighthGnom/http-rest-api/models"
	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		user    func() *models.User
		isValid bool
	}{
		{
			name: "no email",
			user: func() *models.User {
				user := models.TestUser(t)
				user.Email = ""
				return user
			},
			isValid: false,
		},

		{
			name: "invalid email",
			user: func() *models.User {
				user := models.TestUser(t)
				user.Email = "testusercom"
				return user
			},
			isValid: false,
		},

		{
			name: "no password",
			user: func() *models.User {
				user := models.TestUser(t)
				user.Password = ""
				return user
			},
			isValid: false,
		},

		{
			name: "invalid password",
			user: func() *models.User {
				user := models.TestUser(t)
				user.Password = "qwe"
				return user
			},
			isValid: false,
		},

		{
			name: "valid user",
			user: func() *models.User {
				user := models.TestUser(t)
				return user
			},
			isValid: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.isValid {
				assert.NoError(t, testCase.user().Validate())
			} else {
				assert.Error(t, testCase.user().Validate())
			}
		})
	}
}

func TestUser_EnscriptPassword(t *testing.T) {
	user := models.TestUser(t)
	err := user.EnscriptPassword()
	assert.NoError(t, err)
	assert.NotEmpty(t, user.EnscriptedPassword)

}
