package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/eighthGnom/http-rest-api/internal/app/storage/postgresstorage"
	"github.com/gorilla/sessions"
)

func Start(config *Config) error {
	db, err := newBD(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := postgresstorage.New(db)
	sessionsStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionsStore, config.LoggerLevel)

	srv.logger.Infof("Starting server at port: %s", config.BindAddr)
	return http.ListenAndServe(config.BindAddr, srv)
}

func newBD(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}
