package postgresstorage

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

func TestStorage(t *testing.T, dbURL string) (*Storage, func(...string)) {
	t.Helper()
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		t.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}
	storage := New(db)
	if err != nil {
		t.Fatal(err)
	}
	return storage, func(tables ...string) {
		_, err := storage.db.Exec(fmt.Sprintf("truncate %s cascade", strings.Join(tables, ", ")))
		if err != nil {
			t.Fatal(err)
		}
		storage.db.Close()
	}
}
