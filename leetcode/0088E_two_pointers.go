package main

func merge88(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1    // last valid index in nums1
	p2 := n - 1    // last index in nums2
	k := m + n - 1 // last index in nums1 (write position)

	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[k] = nums1[p1]
			p1--
		} else {
			nums1[k] = nums2[p2]
			p2--
		}

		k--
	}
}
