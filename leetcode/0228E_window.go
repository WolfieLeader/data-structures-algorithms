package main

import (
	"strconv"
	"strings"
)

// TODO:

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var sb strings.Builder
	sb.Grow(4) // "x->y"

	out := make([]string, 0, len(nums))

	start := 0
	for start < len(nums) {
		end := start

		// NOTE: LOOP until the next value of nums[end] isn't in range
		for (end+1 < len(nums)) && (nums[end+1]-1 == nums[end]) {
			end++
		}

		sb.Reset()
		sb.WriteString(strconv.Itoa(nums[start]))
		if start != end {
			sb.WriteString("->")
			sb.WriteString(strconv.Itoa(nums[end]))
		}
		out = append(out, sb.String())

		start = end + 1
	}

	return out
}
