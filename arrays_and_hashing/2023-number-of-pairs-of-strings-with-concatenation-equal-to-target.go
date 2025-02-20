/*
Given an array of digit strings `nums` and a digit string `target`, return the number of pairs of indices `(i, j)` where `i != j` such that the
concatenation of `nums[i] + nums[j]` equals `target`.

Example 1:
	Input: nums = ["777","7","77","77"], target = "7777"
	Output: 4
	Explanation: Valid pairs are:
		- (0, 1): "777" + "7"
		- (1, 0): "7" + "777"
		- (2, 3): "77" + "77"
		- (3, 2): "77" + "77"

Example 2:
	Input: nums = ["123","4","12","34"], target = "1234"
	Output: 2
	Explanation: Valid pairs are:
		- (0, 1): "123" + "4"
		- (2, 3): "12" + "34"

Example 3:
	Input: nums = ["1","1","1"], target = "11"
	Output: 6
	Explanation: Valid pairs are:
		- (0, 1): "1" + "1"
		- (1, 0): "1" + "1"
		- (0, 2): "1" + "1"
		- (2, 0): "1" + "1"
		- (1, 2): "1" + "1"
		- (2, 1): "1" + "1"
*/

package arrays_and_hashing

func NumOfPairs(nums []string, target string) int {
	if len(nums) < 2 {
		return 0
	}

	counter, m := 0, map[string]int{}

	for _, str := range nums {
		m[str]++
	}

	for i := 1; i < len(target); i++ {
		a, b := target[:i], target[i:]
		if a != b {
			counter += m[a] * m[b]
			// continue
		}

		counter += m[a] * (m[a] - 1)
	}

	return counter
}
