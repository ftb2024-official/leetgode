package yandex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name     string
		input1   *TreeNode
		input2   *TreeNode
		expected bool
	}{
		{
			name:     "first case",
			input1:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			input2:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			expected: true,
		},
		{
			name:     "second case",
			input1:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			input2:   &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 3}},
			expected: false,
		},
		{
			name:     "third case",
			input1:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
			input2:   &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := isSameTree(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, result)
		})
	}
}
