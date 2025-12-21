package main

// HACK: Core Pattern of Slow Fast Pointers in A Linked List

// Core Pattern of Fast and Slow Pointers
//
// Pattern
// Use two pointers moving at different speeds.
// Slow moves one step.
// Fast moves two steps.
//
// Key idea
// When the fast pointer reaches the end,
// the slow pointer is at the middle.
//
// Invariant
// Fast has traveled twice the distance of slow.
//
// Why it works
// Relative speed guarantees slow reaches the midpoint
// without counting the list length.
// Uses constant space and one pass.

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
