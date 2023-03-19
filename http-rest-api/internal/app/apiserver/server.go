package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/eighthGnom/http-rest-api/internal/app/models"
	"github.com/eighthGnom/http-rest-api/internal/app/storage"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

const (
	sessionName = "mishashttpserver"
)

type server struct {
	logger        *logrus.Logger
	router        *mux.Router
	store         storage.Storage
	sessionsStore sessions.Store
}

func newServer(store storage.Storage, sessionStore sessions.Store, loggerLevel ...string) *server {
	srv := &server{
		logger:        logrus.New(),
		router:        mux.NewRouter(),
		store:         store,
		sessionsStore: sessionStore,
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
	s.router.HandleFunc("/healthcheck", s.healthCheck()).Methods("GET")
	s.router.HandleFunc("/users", s.handleUserCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")
}

func (s *server) healthCheck() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Okay"))
	}
}

func (s *server) handleUserCreate() http.HandlerFunc {
	type requestHelper struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		request := &requestHelper{}
		err := json.NewDecoder(r.Body).Decode(request)
		if err != nil {
			s.errorMessage(w, r, http.StatusBadRequest, err)
			return
		}

		userTest, _ := s.store.User().FindUserByEmail(request.Email)
		if userTest != nil {
			s.errorMessage(w, r, http.StatusConflict, storage.ErrRecordAlreadyExist)
			return
		}

		user := &models.User{
			Email:    request.Email,
			Password: request.Password,
		}
		err = s.store.User().Create(user)
		if err != nil {
			s.errorMessage(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		user.Sanitize()
		s.respond(w, r, http.StatusCreated, user)
	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {
	type requestHelper struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		request := &requestHelper{}
		err := json.NewDecoder(r.Body).Decode(request)
		if err != nil {
			s.errorMessage(w, r, http.StatusBadRequest, err)
			return
		}
		user, err := s.store.User().FindUserByEmail(request.Email)
		if err != nil || !user.CompareEnscriptedPassword(request.Password) {
			s.errorMessage(w, r, http.StatusUnauthorized, storage.ErrEmailOrPasswordInvalid)
			return
		}

		session, err := s.sessionsStore.Get(r, sessionName)
		if err != nil {
			s.errorMessage(w, r, http.StatusInternalServerError, err)
			return
		}
		session.Values["user_id"] = user.ID

		err = s.sessionsStore.Save(r, w, session)
		if err != nil {
			s.errorMessage(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)

	}

}

func (s *server) errorMessage(writer http.ResponseWriter, request *http.Request, statusCode int, err error) {
	s.respond(writer, request, statusCode, map[string]string{
		"error": err.Error(),
	})
}

func (s *server) respond(writer http.ResponseWriter, request *http.Request, statusCode int, data interface{}) {
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(data)
}
