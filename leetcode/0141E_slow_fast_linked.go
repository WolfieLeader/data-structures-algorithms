package main

// HACK: Core Pattern of Slow Fast Pointers in A Linked List

// Slow and Fast Pointers in Linked List
//
// Pattern
// Use two pointers that move at different speeds.
// Slow moves one step at a time.
// Fast moves two steps at a time.
//
// Invariant
// If there is no cycle, the fast pointer will reach nil.
// If there is a cycle, the fast pointer will eventually
// meet the slow pointer inside the cycle.
//
// Pointer movement
// Always move pointers first, then compare.
// Equality at any point after movement means a cycle exists.
//
// Why it works
// The fast pointer gains one node per step inside the cycle.
// This guarantees they must meet if a cycle exists.
// Each pointer moves at most O(n) steps.
//
// Walkthrough
// slow and fast start at head
// slow moves 1 step
// fast moves 2 steps
// repeat until fast is nil or slow == fast

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
