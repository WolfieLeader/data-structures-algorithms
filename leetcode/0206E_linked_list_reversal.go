package main

// HACK: Core Pattern for Linked List Reversal

// Linked List Reversal
//
// Pattern
// Reverse pointers one node at a time while traversing the list.
//
// Key idea
// For each node, redirect its Next pointer to the previous node.
// This flips the direction of the list in place.
//
// Invariant
// prev always points to the head of the already reversed part.
// curr points to the node currently being processed.
// The rest of the list starting at curr is still unreversed.
//
// Why it works
// Each node is visited once and its pointer is reversed.
// No extra data structures are needed.
// Time complexity is linear with constant extra space.
//
// Walkthrough
// prev = nil, curr = head
// save curr.Next
// point curr.Next to prev
// move prev and curr forward
// repeat until curr is nil
// prev becomes the new head

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next // Store to move forward

		curr.Next = prev // Switch pointer for current node
		prev = curr      // Advance the prev to current node

		curr = next // Move forward
	}

	// `prev` is the new head, `curr` is nil
	return prev
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
