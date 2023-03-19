package apiserver

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/eighthGnom/http-rest-api/internal/app/models"
	"github.com/eighthGnom/http-rest-api/internal/app/storage"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

const (
	sessionName        = "mishashttpserver"
	userKey     ctxKey = iota
	uuidKey     ctxKey = iota
)

type ctxKey int

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
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequst)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/healthcheck", s.healthCheck()).Methods("GET")
	s.router.HandleFunc("/users", s.handleUserCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")

	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authUser)
	private.HandleFunc("/whuemi", s.whuEmI())
}

func (s *server) healthCheck() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Okay"))
	}
}

func (s *server) whuEmI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(userKey))
	}
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), uuidKey, id)))
	})

}

func (s *server) logRequst(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote addr": r.RemoteAddr,
			"request id":  r.Context().Value(uuidKey),
		})
		start := time.Now()
		logger.Infof("started %s %s", r.Method, r.RequestURI)
		rw := &responseWriter{
			ResponseWriter: w,
			code:           http.StatusTeapot,
		}
		next.ServeHTTP(rw, r)
		stop := time.Since(start)
		logger.Infof("finished in %s whit code %v %s", stop, rw.code, http.StatusText(rw.code))

	})
}

func (s *server) authUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionsStore.Get(r, sessionName)
		if err != nil {
			s.errorMessage(w, r, http.StatusInternalServerError, err)
			return
		}
		if session.IsNew {
			s.errorMessage(w, r, http.StatusUnauthorized, storage.ErrUserUnauthorized)
			return
		}
		id, ok := session.Values["user_id"]
		if !ok {
			s.errorMessage(w, r, http.StatusUnauthorized, storage.ErrUserUnauthorized)
			return
		}
		user, err := s.store.User().FindUserByID(id.(int))
		if err != nil {
			s.errorMessage(w, r, http.StatusUnauthorized, storage.ErrUserUnauthorized)
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userKey, user)))
	})

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
