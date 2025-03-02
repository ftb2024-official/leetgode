package yandex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "first case",
			input:    "()",
			expected: true,
		},
		{
			name:     "2-nd case",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "3-rd case",
			input:    "(]",
			expected: false,
		},
		{
			name:     "4-th case",
			input:    "([])",
			expected: true,
		},
		{
			name:     "5-th case",
			input:    "}])[]([{",
			expected: false, // don't know why 'false'
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := isValid(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
