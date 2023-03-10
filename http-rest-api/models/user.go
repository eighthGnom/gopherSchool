package models

type User struct {
	ID                 int    `json:"id"`
	Email              string `json:"email"`
	EnscriptedPassword string `json:"enscripted_password"`
}
