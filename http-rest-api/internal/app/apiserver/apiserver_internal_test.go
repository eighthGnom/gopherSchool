package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eighthGnom/http-rest-api/storage/teststorage"
	"github.com/stretchr/testify/assert"
)

func TestServer_HealthCheck(t *testing.T) {
	store := teststorage.New()
	srv := newServer(store)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/healthcheck", nil)
	srv.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Okay", string(recorder.Body.Bytes()))
}
