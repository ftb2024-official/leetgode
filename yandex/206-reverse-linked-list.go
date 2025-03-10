/*
Given the `head` of a singly linked list, reverse the list, and return the reversed list.

Example 1:
	Input: head = [1,2,3,4,5]
	Output: [5,4,3,2,1]

Example 2:
	Input: head = [1,2]
	Output: [2,1]

Example 3:
	Input: head = []
	Output: []
*/

package yandex

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head

	for current != nil {
		nexTemp := current.Next
		current.Next = prev
		prev = current
		current = nexTemp
	}

	return prev
}
