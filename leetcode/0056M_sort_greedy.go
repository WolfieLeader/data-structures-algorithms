package main

import "sort"

// TODO:

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		curr := intervals[i]

		// Last end is not bigger than current start - Overlap (touching counts)
		if last[1] >= curr[0] {
			last[1] = max(last[1], curr[1])
		} else {
			merged = append(merged, curr)
		}
	}

	return merged
}
