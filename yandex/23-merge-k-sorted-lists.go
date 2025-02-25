/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order. Merge all the linked-lists into one sorted
linked-list and return it.

Example 1:
	Input: lists = [[1,4,5],[1,3,4],[2,6]]
	Output: [1,1,2,3,4,4,5,6]
	Explanation: The linked-lists are:
	[
		1->4->5,
		1->3->4,
		2->6
	]
	merging them into one sorted list: 1->1->2->3->4->4->5->6

Example 2:
	Input: lists = []
	Output: []

Example 3:
	Input: lists = [[]]
	Output: []
*/

package yandex

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

// MinHeap Approach
type MinHeap []*ListNode

func (mh MinHeap) Len() int           { return len(mh) }
func (mh MinHeap) Less(i, j int) bool { return mh[i].Val < mh[j].Val }
func (mh MinHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }

func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.(*ListNode))
}

func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	last := old[n-1]
	*mh = old[:n-1]
	return last
}

func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}

	dummy := &ListNode{}
	current := dummy

	for minHeap.Len() > 0 {
		// extract the smallest node
		smallest := heap.Pop(minHeap).(*ListNode)
		current.Next = smallest
		current = current.Next

		// if there is a next node, push it into the heap
		if smallest.Next != nil {
			heap.Push(minHeap, smallest.Next)
		}
	}

	return dummy.Next
}

func Example1() {
	list1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	list3 := &ListNode{2, &ListNode{6, nil}}

	lists := []*ListNode{list1, list2, list3}

	mergedList := mergeKLists(lists)
	printList(mergedList)
}

// Divide & Conquer Approach

// merge two sorted linked lists (same as in merge sort)
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}

		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	}

	if l2 != nil {
		current.Next = l2
	}

	return dummy.Next
}

// recursive function to merge lists in a divide & conquer approach
func mergeLists(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	mid := (left + right) / 2
	l1 := mergeLists(lists, left, mid)
	l2 := mergeLists(lists, mid+1, right)

	return mergeTwoLists(l1, l2)
}

func mergeKLists_2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return mergeLists(lists, 0, len(lists)-1)
}

func Example2() {
	list1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	list3 := &ListNode{2, &ListNode{6, nil}}

	lists := []*ListNode{list1, list2, list3}

	mergedList := mergeKLists_2(lists)
	printList(mergedList)
}
