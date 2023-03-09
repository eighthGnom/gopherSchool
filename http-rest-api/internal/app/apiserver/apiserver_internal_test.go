package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	server := New(NewConfig())
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/health-check", nil)
	server.healthCheck().ServeHTTP(recorder, request)
	assert.Equal(t, "Okay", recorder.Body.String())
}
