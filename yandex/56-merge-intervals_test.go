package yandex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name:     "first case",
			input:    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:     "second case",
			input:    [][]int{{1, 4}, {4, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			name:     "third case",
			input:    [][]int{{5, 7}, {7, 9}, {-1, 2}, {3, 4}, {-5, 1}},
			expected: [][]int{{-5, 2}, {3, 4}, {5, 9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := merge(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
