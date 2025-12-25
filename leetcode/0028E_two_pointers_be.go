package main

// TODO:

// I: haystack = "sadbutsad", needle = "sad"  =>  O: 0
// I: haystack = "leetcode", needle = "leeto" =>  O: -1

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	left := 0
	for left <= len(haystack)-len(needle) {
		right := 0

		for right < len(needle) && haystack[left+right] == needle[right] {
			right++
		}

		if right == len(needle) {
			return left
		}

		left++
	}

	return -1
}
