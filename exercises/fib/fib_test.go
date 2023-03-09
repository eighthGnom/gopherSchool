package fib_test

import (
	"gopherSchool/exercises/fib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "zero",
			input:    0,
			expected: 0,
		},
		{
			name:     "one",
			input:    1,
			expected: 1,
		},
		{
			name:     "two",
			input:    2,
			expected: 1,
		},
		{
			name:     "three",
			input:    3,
			expected: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, fib.Fib(testCase.input))
		})
	}
}
