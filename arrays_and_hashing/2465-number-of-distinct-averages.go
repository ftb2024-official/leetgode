/*
You are given a 0-indexed integer array `nums` of even length. As long as `nums` is not empty, you must repetitively: 1) Find the minimum number in
`nums` and remove it; 2) Find the maximum number in `nums` and remove it; 3) Calculate the average of the two removed numbers. Return the number of
distinct averages calculated using the above process. Note that when there is a tie for a minimum or maximum number, any can be removed.

Example 1:
	Input: nums = [4,1,4,0,3,5]
	Output: 2
	Explanation:
		1. Remove 0 and 5, and the average is (0 + 5) / 2 = 2.5. Now, nums = [4,1,4,3].
		2. Remove 1 and 4. The average is (1 + 4) / 2 = 2.5, and nums = [4,3].
		3. Remove 3 and 4, and the average is (3 + 4) / 2 = 3.5.
		Since there are 2 distinct numbers among 2.5, 2.5, and 3.5, we return 2.

Example 2:
	Input: nums = [1,100]
	Output: 1
	Explanation:
		There is only one average to be calculated after removing 1 and 100, so we return 1.
*/

package arrays_and_hashing

import (
	"fmt"
	"slices"
)

func DistinctAverages(nums []int) int {
	slices.Sort(nums)
	i, j, avg, m := 0, len(nums)-1, 0.0, map[float64]struct{}{}

	fmt.Println(nums)

	for i < j {
		avg = (float64(nums[i]+nums[j]) / 2.0)
		if _, ok := m[avg]; !ok {
			m[avg] = struct{}{}
		}

		i++
		j--
	}

	fmt.Println(m)

	return len(m)
}
