package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums     []int
	target   int
	expected []int
}

var cases = []Case{
	{
		nums:     []int{2, 7, 11, 15},
		target:   9,
		expected: []int{0, 1},
	},
	{
		nums:     []int{3, 2, 4},
		target:   6,
		expected: []int{1, 2},
	},
	{
		nums:     []int{3, 3},
		target:   6,
		expected: []int{0, 1},
	},
}

func TestTwoSum(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, twoSum(c.nums, c.target), c.expected)
	}
}
