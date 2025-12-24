package main

// TODO:

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	fast, slow := dummy, dummy

	// Move fast n steps ahead
	for range n {
		fast = fast.Next
	}

	// Move both pointers until fast is at the end
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Remove nth node from end
	slow.Next = slow.Next.Next

	return dummy.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
