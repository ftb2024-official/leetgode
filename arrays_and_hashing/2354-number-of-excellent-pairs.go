/*
You are given a 0-indexed positive integer array `nums` and a positive integer `k`. A pair of numbers `(num1, num2)` is called excellent is the
following conditions are satisfied: 1) Both the numbers `num1` and `num2` exist in the array `nums`. 2) The sum of the number of set bits in
`num1 OR num2` and `num1 AND num2` is greater than or equl to `k` where OR is the bitwise OR operation and AND is the bitwise AND operation. Return
the number of distinct excellent pairs. Two pairs `(a, b)` and `(c, d)` are considered distinct if either `a != c` or `b != d`. For example
`(1, 2)` and `(2, 1)` are distinct. Note that a pair `(num1, num2)` such that `num1 == num2` can also be excellent if you have at least one
occurrence of `num1` in the array.

Example 1:
	Input: nums = [1,2,3,1], k = 3
	Output: 5
	Explanation: The excellent pairs are the following:
		- (3, 3). (3 AND 3) and (3 OR 3) are both equal to (11) in binary. The total number of set bits is 2 + 2 = 4, which is greater than or
		equal to k = 3.
		- (2, 3) and (3, 2). (2 AND 3) is equal to (10) in binary, and (2 OR 3) is equal to (11) in binary. The total number of set bits is 1 + 2 =
		3.
		- (1, 3) and (3, 1). (1 AND 3) is equal to (01) in binary, and (1 OR 3) is equal to (11) in binary. The total number of set bits is 1 + 2 =
		3.
		So the number of excellent pairs is 5.

Example 2:
	Input: nums = [5,1,1], k = 10
	Output: 0
	Explanation: There are no excellent pairs for this array.
*/

package arrays_and_hashing

import (
	"math/bits"
	"sort"
)

func CountExcellentPairs(nums []int, k int) int {
	// Count the number of bits in each unique number
	bitCounts := make(map[int]int)
	for _, num := range nums {
		bitCounts[bits.OnesCount(uint(num))]++
	}

	// Sort the (bit count, tally) pairs by bit count
	var counts [][]int
	for count, tally := range bitCounts {
		counts = append(counts, []int{count, tally})
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i][0] < counts[j][0]
	})

	// Compute the reversed prefix sum of the tallies
	sums := make([]int, len(counts))
	sums[len(sums)-1] = counts[len(counts)-1][1]
	for i := len(sums) - 2; i >= 0; i-- {
		sums[i] = counts[i][1] + sums[i+1]
	}

	// Choose each number as the first number of a pair
	pairs := 0
	for _, pair := range counts {
		n, c := pair[0], pair[1]

		// Find the smallest number which forms a valid pair
		i := sort.Search(len(counts), func(idx int) bool {
			return counts[idx][0] >= k-n
		})

		// Check if any pairs can be formed
		if i < len(sums) {
			// Compute the number of pairs which start with the given collection of numbers
			pairs += c * sums[i]
		}
	}

	// Return the number of pairs
	return pairs
}
