package main

// Core Pattern of Floyd Cycle Start
//
// Phase 1
// Use slow and fast pointers to detect a cycle.
// If they meet, a cycle exists.
//
// Phase 2
// Set one pointer to head.
// Keep the other at the meeting point.
// Move both one step at a time.
// The node where they meet is the cycle start.
//
// If no cycle, return nil.

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// Meet occurred
		if slow == fast {
			p1 := head
			p2 := slow

			// NOTE: LOOP until p1 and p2 meet again
			for p1 != p2 {
				p1 = p1.Next
				p2 = p2.Next
			}

			return p1
		}
	}

	return nil
}
