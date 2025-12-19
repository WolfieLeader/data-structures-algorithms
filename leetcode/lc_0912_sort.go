package main

func optimizedSortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	if min == max {
		return nums
	}

	count := make([]int, max-min+1)
	for _, num := range nums {
		count[num-min]++
	}

	i := 0
	for j := 0; j < len(count); j++ {
		for count[j] > 0 {
			nums[i] = j + min
			count[j]--
			i++
		}
	}

	return nums
}

func sortArray(nums []int) []int {
	if len(nums) > 1 {
		quickSort(nums, 0, len(nums)-1)
	}
	return nums
}

func quickSort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	lt, gt := partition3Way(nums, lo, hi)
	quickSort(nums, lo, lt-1)
	quickSort(nums, gt+1, hi)
}

func partition3Way(nums []int, lo, hi int) (int, int) {
	pivot := nums[lo+(hi-lo)/2]

	i, lt, gt := lo, lo, hi
	for i <= gt {
		switch {
		case nums[i] < pivot:
			nums[lt], nums[i] = nums[i], nums[lt] //Swap
			lt++
			i++
		case nums[i] > pivot:
			nums[gt], nums[i] = nums[i], nums[gt] //Swap
			gt--
		default:
			i++
		}
	}

	return lt, gt
}
