package apiserver

import (
	"net/http"

	"github.com/eighthGnom/http-rest-api/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	logger *logrus.Logger
	router *mux.Router
	store  storage.Storage
}

func newServer(store storage.Storage, loggerLevel ...string) *server {
	srv := &server{
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store,
	}
	if len(loggerLevel) > 0 {
		if err := srv.configureLogger(loggerLevel[0]); err != nil {
			logrus.Error("cant configure logger field", err)
		}
	}
	srv.configureRouter()

	return srv
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureLogger(logLevel string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *server) configureRouter() {
	s.router.Handle("/healthcheck", s.healthCheck()).Methods("GET")
}

func (s *server) healthCheck() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Okay"))
	}
}
