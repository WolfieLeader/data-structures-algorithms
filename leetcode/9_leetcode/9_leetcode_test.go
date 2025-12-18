package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	input    int
	expected bool
}

var cases = []Case{
	{input: 121, expected: true},
	{input: -121, expected: false},
	{input: 10, expected: false},
}

func TestIsPalindrome(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.expected, isPalindrome(c.input), "isPalindrome(%d)", c.input)
	}
}

func TestFasterIsPalindrome(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.expected, fasterIsPalindrome(c.input), "fasterIsPalindrome(%d)", c.input)
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for b.Loop() {
		isPalindrome(12345678987654321)
	}
}

func BenchmarkFasterIsPalindrome(b *testing.B) {
	for b.Loop() {
		fasterIsPalindrome(12345678987654321)
	}
}
