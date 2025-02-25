package patterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixSumArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Basic case",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 3, 6, 10},
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "All zeros",
			input:    []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "Negative numbers",
			input:    []int{-1, -2, -3, -4},
			expected: []int{-1, -3, -6, -10},
		},
		{
			name:     "Mixed case",
			input:    []int{-1, 2, -3, 4},
			expected: []int{-1, 1, -2, 2},
		},
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "nil case",
			input:    nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := prefixSumArray(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
