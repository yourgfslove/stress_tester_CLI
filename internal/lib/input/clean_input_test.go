package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClean(t *testing.T) {
	testcases := []struct {
		name    string
		input   string
		correct []string
	}{{
		name:    "default",
		input:   "hello world",
		correct: []string{"hello", "world"},
	},
		{
			name:    "a lot spaces",
			input:   "hello           world          boy",
			correct: []string{"hello", "world", "boy"},
		},
		{
			name:    "russian language",
			input:   "привет мир",
			correct: []string{"привет", "мир"},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			input := Clean(tc.input)
			assert.Equal(t, tc.correct, input)
		})
	}
}
