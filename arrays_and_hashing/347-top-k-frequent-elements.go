/*
Given an integer array `nums` and an integer `k`, return the `k` most frequent elements. You may return the answer in any order.

Example 1:
	Input: nums = [1,1,1,2,2,3], k = 2
	Output: [1,2]

Example 2:
	Input: nums = [1], k = 1
	Output: [1]
*/

package arrays_and_hashing

import (
	"container/heap"
	"sort"
)

func TopKFrequent(nums []int, k int) []int {
	ans, frequency := make([]int, k), make(map[int]int)

	for _, num := range nums {
		frequency[num]++
	}

	// convert the frequency map to an array of pairs
	var pairs [][]int
	for num, freq := range frequency {
		pairs = append(pairs, []int{num, freq})
	}

	// sort the pairs on descending order based on frequency
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	// extract the first k elements from the sorted array
	for i := 0; i < k; i++ {
		ans[i] = pairs[i][0]
	}

	return ans
}

type Element struct {
	val  int
	freq int
}

type MinHeap []*Element

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].freq < h[j].freq }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*Element)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func TopKFrequent_2(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// create a min-heap
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// add k elements to the min-heap
	for num, freq := range freq {
		if minHeap.Len() < k {
			heap.Push(minHeap, &Element{val: num, freq: freq})
		} else if freq > (*minHeap)[0].freq {
			heap.Pop(minHeap)
			heap.Push(minHeap, &Element{val: num, freq: freq})
		}
	}

	// extract elements from min-heap
	ans := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(minHeap).(*Element).val
	}

	return ans
}

/*
In Go, `type MinHeap []*Element` and `type MinHeap *[]Element` represent different types of min heaps, specifically slices of pointers to `Element` and a pointer to a slice of `Element`, respectively. Here's the breakdown:

1. **`type MinHeap []*Element`**:
   - This declaration defines `MinHeap` as a new type, which is a slice of pointers to `Element`.
   - Each element of the slice is a pointer to an `Element` struct.
   - This is commonly used when implementing a heap data structure where each element in the heap is represented by a pointer, and the heap itself is maintained as a slice of pointers.

2. **`type MinHeap *[]Element`**:
   - This declaration defines `MinHeap` as a new type, which is a pointer to a slice of `Element`.
   - The slice contains elements of type `Element`.
   - This is less common in heap implementations but could be used if you want to pass around a pointer to a slice directly instead of creating a new type.

In summary, the main difference lies in how the slice is managed and accessed. With `MinHeap []*Element`, you're working directly with a slice
of pointers to `Element`, while with `MinHeap *[]Element`, you're working with a pointer to a slice of `Element`. Depending on your use case and
preferences, you may choose one over the other. The former is more common in heap implementations in Go.
*/
