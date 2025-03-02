package yandex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		input1   []int
		input2   int
		expected int
	}{
		{
			name:     "first case",
			input1:   []int{-1, 0, 3, 5, 9, 12},
			input2:   9,
			expected: 4,
		},
		{
			name:     "second case",
			input1:   []int{-1, 0, 3, 5, 9, 12},
			input2:   2,
			expected: -1,
		},
		{
			name:     "one element case with -1",
			input1:   []int{-1},
			input2:   2,
			expected: -1,
		},
		{
			name:     "one element case without -1",
			input1:   []int{-1},
			input2:   -1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := binarySearch(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, result)
		})
	}
}
