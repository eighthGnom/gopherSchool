package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpFib(t *testing.T) {
	tesstCases := []struct {
		name     string
		input    int
		expected []byte
	}{
		{
			name:     "zero",
			input:    0,
			expected: []byte("0"),
		},

		{
			name:     "one",
			input:    1,
			expected: []byte("1"),
		},

		{
			name:     "two",
			input:    2,
			expected: []byte("1"),
		},

		{
			name:     "three",
			input:    3,
			expected: []byte("2"),
		},
	}
	for _, testCase := range tesstCases {
		t.Run(testCase.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/fibonachi?num=%d", testCase.input), nil)
			HttpFib().ServeHTTP(recorder, request)
			assert.Equal(t, testCase.expected, recorder.Body.Bytes())
		})
	}
}
