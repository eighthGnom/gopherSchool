package models

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()
	return &User{
		Email:    "testuser@gmail.com",
		Password: "Qwerty999)))",
	}
}
