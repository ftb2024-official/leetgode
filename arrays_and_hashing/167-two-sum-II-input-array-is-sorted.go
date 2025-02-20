/*
Given a 1-indexed array of integers `numbers` that is already sorted in ascending order, find two numbers such that they add up to specific
`target` number. Let these two numbers be `numbers[idx1]` and `numbers[idx2]`, where `1 <= idx1 < idx2 <= numbers.length`. Return the indices of
the two numbers `idx1` and `idx2`, added by one as an integer array `[idx1, idx2]` of length 2. The tests are generated such that there is exacyly
one solution. You may not use the same element twice. Your solution must use only constant extra space.

Example 1:
	Input: numbers = [2,7,11,15], target = 9
	Output: [1,2]
	Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

Example 2:
	Input: numbers = [2,3,4], target = 6
	Output: [1,3]
	Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

Example 3:
	Input: numbers = [-1,0], target = -1
	Output: [1,2]
	Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
*/

package arrays_and_hashing

func TwoSumII(numbers []int, target int) []int {
	result := []int{}

	for i := 0; i < len(numbers)-1; i++ {
		if i > 0 && numbers[i] == numbers[i-1] {
			continue
		}

		right := len(numbers) - 1

		for i < right {
			twoSum := numbers[i] + numbers[right]
			if twoSum > target {
				right--
			} else if twoSum < target {
				i++
			} else {
				result = append(result, i+1, right+1)

				i++
				right--
			}
		}
	}

	return result
}
