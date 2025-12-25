package main

func intersect(nums1, nums2 []int) []int {
	// Count the smaller array to save memory
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	counts := make(map[int]int, len(nums1))
	for _, x := range nums1 {
		counts[x]++
	}

	out := make([]int, 0)
	for _, x := range nums2 {
		if counts[x] > 0 {
			out = append(out, x)
			counts[x]--
		}
	}
	return out
}
