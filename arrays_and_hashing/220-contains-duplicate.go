/*
You are given an integer array `nums` and two integers `indexDiff` and `valueDiff`. Find a pair of indices `(i, j)` such that:
	- `i != j`
	- `abs(i - j) <= indexDiff`
	- `abs(nums[i] - nums[j]) <= valueDiff`.
Return `true` if such pair exists otherwise `false`.

Example 1:
	Input: nums = [1,2,3,1], indexDiff = 3, valueDiff = 0
	Output: true
	Explanation: We can choose (i, j) = (0, 3). We satisfy the three conditions: i != j --> 0 != 3 abs(i - j) <= indexDiff --> abs(0 - 3) <= 3
	abs(nums[i] - nums[j]) <= valueDiff --> abs(1 - 1) <= 0

	Example 2:
	Input: nums = [1,5,9,1,5,9], indexDiff = 2, valueDiff = 3
	Output: false
	Explanation: After trying all the possible pairs (i, j), we cannot satisfy the three conditions, so we return false.

Constraints:
	2 <= nums.length <= 10^5
	-10^9 <= nums[i] <= 10^9
	1 <= indexDiff <= nums.length
	0 <= valueDiff <= 10^9
*/

package arrays_and_hashing

import "math"

func ContainsNearbyAlmostDuplicate(nums []int, indexDiff, valueDiff int) bool {
	if (len(nums) <= 1 || len(nums) >= int(math.Pow10(5))) ||
		(indexDiff < 1 || indexDiff > len(nums)) ||
		(valueDiff < 0 || valueDiff > int(math.Pow10(9))) {
		return false
	}

	seen := make(map[int]int)

	for idx, num := range nums {
		bucket := num / (valueDiff + 1)

		if _, exists := seen[bucket]; exists {
			return true
		}

		if val, exists := seen[bucket-1]; exists && math.Abs(float64(num-val)) <= float64(valueDiff) {
			return true
		}

		if val, exists := seen[bucket+1]; exists && math.Abs(float64(num-val)) <= float64(valueDiff) {
			return true
		}

		seen[bucket] = num

		if idx >= indexDiff {
			delete(seen, nums[idx-indexDiff]/(valueDiff+1))
		}
	}

	return false
}
