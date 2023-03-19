package models_test

import (
	"testing"

	models2 "github.com/eighthGnom/http-rest-api/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		user    func() *models2.User
		isValid bool
	}{
		{
			name: "no email",
			user: func() *models2.User {
				user := models2.TestUser(t)
				user.Email = ""
				return user
			},
			isValid: false,
		},

		{
			name: "invalid email",
			user: func() *models2.User {
				user := models2.TestUser(t)
				user.Email = "testusercom"
				return user
			},
			isValid: false,
		},

		{
			name: "no password",
			user: func() *models2.User {
				user := models2.TestUser(t)
				user.Password = ""
				return user
			},
			isValid: false,
		},

		{
			name: "invalid password",
			user: func() *models2.User {
				user := models2.TestUser(t)
				user.Password = "qwe"
				return user
			},
			isValid: false,
		},

		{
			name: "valid user",
			user: func() *models2.User {
				user := models2.TestUser(t)
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
	user := models2.TestUser(t)
	err := user.EnscriptPassword()
	assert.NoError(t, err)
	assert.NotEmpty(t, user.EnscriptedPassword)

}
