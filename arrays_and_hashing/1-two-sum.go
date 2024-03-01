/*
Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`. You may assume that each
input would have exactly one solution, and you may not use the same element twice. You can return an answer in any order.

Example 1:
	Input: nums = [2,7,11,15], target = 9
	Output: [0,1]
	Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
	Input: nums = [3,2,4], target = 6
	Output: [1,2]

Example 3:
	Input: nums = [3,3], target = 6
	Output: [0,1]
*/

package arraysandhashing

func TwoSum(nums []int, target int) []int {
	kv := make(map[int]int)

	for idx, num := range nums {
		sub := target - num
		if idx2, ok := kv[sub]; ok {
			return []int{idx, idx2}
		} else {
			kv[num] = idx
		}
	}

	return nil
}
