package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eighthGnom/http-rest-api/internal/app/models"
	"github.com/eighthGnom/http-rest-api/internal/app/storage/teststorage"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestServer_HealthCheck(t *testing.T) {
	store := teststorage.New()
	sessionsStore := sessions.NewCookieStore([]byte("session"))
	srv := newServer(store, sessionsStore)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/healthcheck", nil)
	srv.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Okay", string(recorder.Body.Bytes()))
}

func TestServer_HandleUserCreate(t *testing.T) {
	storage := teststorage.New()
	storage.User().Create(models.TestUser(t))
	sessionsStore := sessions.NewCookieStore([]byte("session"))
	srv := newServer(storage, sessionsStore)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "invalid user",
			payload: map[string]string{
				"email": "invalid",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "valid user",
			payload: map[string]string{
				"email":    "user@gmail.com",
				"password": "Qwerty999)))",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "user exists",
			payload: map[string]string{
				"email":    "testuser@gmail.com",
				"password": "Qwerty999)))",
			},
			expectedCode: http.StatusConflict,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := bytes.Buffer{}
			json.NewEncoder(&b).Encode(tc.payload)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("POST", "/users", &b)
			srv.handleUserCreate().ServeHTTP(recorder, request)
			assert.Equal(t, tc.expectedCode, recorder.Code)

		})
	}
}
func TestServer_HandleSessionsCreate(t *testing.T) {
	store := teststorage.New()
	sessionsStore := sessions.NewCookieStore([]byte("secret"))
	srv := newServer(store, sessionsStore)
	user := models.TestUser(t)
	srv.store.User().Create(user)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "invalid email",
			payload: map[string]interface{}{
				"email":    "wewerwrwer",
				"password": user.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]interface{}{
				"email":    user.Email,
				"password": "wrewerwer",
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "valid user",
			payload: map[string]interface{}{
				"email":    user.Email,
				"password": user.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			json.NewEncoder(buffer).Encode(tc.payload)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("POST", "/sessions", buffer)
			srv.handleSessionsCreate().ServeHTTP(recorder, request)
			assert.Equal(t, tc.expectedCode, recorder.Code)
		})
	}
}
