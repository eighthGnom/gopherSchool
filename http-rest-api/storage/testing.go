package storage

import (
	"fmt"
	"strings"
	"testing"
)

func TestStorage(t *testing.T, dbURL string) (*Storage, func(...string)) {
	t.Helper()
	config := NewConfig()
	config.DatabaseURL = dbURL
	storage := New(config)
	err := storage.Open()
	if err != nil {
		t.Fatal(err)
	}
	return storage, func(tables ...string) {
		_, err := storage.db.Exec(fmt.Sprintf("truncate %s cascade", strings.Join(tables, ", ")))
		if err != nil {
			t.Fatal(err)
		}
		storage.Close()
	}
}
