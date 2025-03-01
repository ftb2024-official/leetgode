package patterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoPointer_SortedArray(t *testing.T) {
	tests := []struct {
		name     string
		input1   []int
		input2   int
		expected []int
	}{
		{
			name:     "basic case",
			input1:   []int{1, 3, 4, 6, 8, 10, 13},
			input2:   13,
			expected: []int{3, 10},
		},
		{
			name:     "single element",
			input1:   []int{1},
			input2:   13,
			expected: []int{},
		},
		{
			name:     "all zeros",
			input1:   []int{0, 0, 0},
			input2:   13,
			expected: []int{},
		},
		{
			name:     "negative numbers",
			input1:   []int{-1, -3, -4, -6, -8, -10, -13},
			input2:   -11,
			expected: []int{-10, -1},
		},
		{
			name:     "mixed case",
			input1:   []int{1, -3, 4, -6, 8, -10, 13},
			input2:   12,
			expected: []int{4, 8},
		},
		{
			name:     "empty slice",
			input1:   []int{},
			input2:   13,
			expected: []int{},
		},
		{
			name:     "nil case",
			input1:   nil,
			input2:   13,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := twoPointer_sorted_array(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTwoPointer_Palindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid palindrome",
			input:    "racecar",
			expected: true,
		},
		{
			name:     "invalid palindrome",
			input:    "something",
			expected: false,
		},
		{
			name:     "single letter word",
			input:    "r",
			expected: true,
		},
		{
			name:     "empty string",
			input:    "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := twoPointer_palindrome(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
