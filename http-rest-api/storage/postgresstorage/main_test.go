package postgresstorage_test

import (
	"os"
	"testing"
)

var databaseURL string

func TestMain(t *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost port=5432 user=postgres dbname=test_gopher_school password=postgres sslmode=disable"
	}
	os.Exit(t.Run())
}
