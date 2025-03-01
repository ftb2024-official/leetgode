package yandex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSum_BruteForce(t *testing.T) {
	tests := []struct {
		name     string
		input1   []int
		input2   int
		expected [][]int
	}{
		{
			name:     "first case",
			input1:   []int{1, 0, -1, 0, -2, 2},
			input2:   0,
			expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := fourSum_brute_force(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, result)
		})
	}
}
