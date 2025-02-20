/*
A good meal is a meal that contains exactly two different food items with a sum of deliciousness equal to a power of two. You can pick any two
different foods to make a good meal. Given an array of integers `deliciousness`, where `deliciousness[i]` is the deliciousness of the `i^th` item
of food, return the number of different good meals you can make from this list modulo `10^9 + 7`. Note that items with different indices are
considered different even if they have the same deliciousness value.

Example 1:
	Input: deliciousness = [1,3,5,7,9]
	Output: 4
	Explanation: The good meals are (1,3), (1,7), (3,5) and, (7,9).
		Their respective sums are 4, 8, 8, and 16, all of which are powers of 2.

Example 2:
	Input: deliciousness = [1,1,1,3,3,3,7]
	Output: 15
	Explanation: The good meals are (1,1) with 3 ways, (1,3) with 9 ways, and (1,7) with 3 ways.
*/

package arrays_and_hashing

func CountPairs(deliciousness []int) int {
	if len(deliciousness) < 1 {
		return 0
	}

	const mod = int(1e9) + 7 // 1 * 10 ^ 9
	counter, deliciousnessMap := 0, map[int]int{}

	for _, dVal := range deliciousness {
		for i := 1; i <= 1<<21; i <<= 1 {
			// if the complement (powerOfTwo - dVal) exists in the map
			if v, ok := deliciousnessMap[i-dVal]; ok {
				counter = (counter + v) % mod
			}
		}

		deliciousnessMap[dVal]++
	}

	return counter
}
