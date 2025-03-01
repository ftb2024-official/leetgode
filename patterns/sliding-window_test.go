package patterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlidingWindow_MaxSumOfSubarrayOfSizeK(t *testing.T) {
	tests := []struct {
		name     string
		input1   []int
		input2   int
		expected int
	}{
		{
			name:     "first case",
			input1:   []int{2, 1, 5, 1, 3, 2},
			input2:   3,
			expected: 9,
		},
		{
			name:     "len less than k",
			input1:   []int{2, 1, 5, 1, 3, 2},
			input2:   7,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := slidingWindow_max_sum_of_subarray_of_size_k(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, result)
		})
	}
}
