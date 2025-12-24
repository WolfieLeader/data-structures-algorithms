package main

// Slow Fast Pointers + Linked List Reversal
//
// Steps
// 1 Find the middle using slow and fast pointers
// 2 Reverse the second half in place
// 3 Compare values from the start and from the reversed half
//
// Why it works
// A palindrome reads the same forward and backward.
// Reversing the second half lets us compare symmetric nodes
// using only constant extra space.

func isPalindrome234(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Odd Length, move to the right of the middle
	if fast != nil {
		slow = slow.Next
	}

	// Reverse
	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next

		curr.Next = prev
		prev = curr

		curr = next
	}

	// Compare halves
	p1, p2 := head, prev // Prev is the new head
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
