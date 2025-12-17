package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	l1       *ListNode
	l2       *ListNode
	expected *ListNode
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []Case{
		{
			l1:       list(2, 4, 3),
			l2:       list(5, 6, 4),
			expected: list(7, 0, 8),
		},
		{
			l1:       list(0),
			l2:       list(0),
			expected: list(0),
		},
		{
			l1:       list(9, 9, 9, 9, 9, 9, 9),
			l2:       list(9, 9, 9, 9),
			expected: list(8, 9, 9, 9, 0, 0, 0, 1),
		},
		{
			l1:       list(1),
			l2:       list(9, 9),
			expected: list(0, 0, 1),
		},
	}

	for i, tc := range cases {
		result := addTwoNumbers(tc.l1, tc.l2)
		assert.Truef(
			t,
			equal(result, tc.expected),
			"case %d failed: got %v, want %v",
			i,
			toSlice(result),
			toSlice(tc.expected),
		)
	}
}

func list(vals ...int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func equal(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}

func toSlice(l *ListNode) []int {
	var out []int
	for l != nil {
		out = append(out, l.Val)
		l = l.Next
	}
	return out
}
