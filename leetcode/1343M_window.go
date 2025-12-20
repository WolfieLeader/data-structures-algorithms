package main

// I: arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4 => Ex: [[2,5,5],[5,5,5],[5,5,8]]: avg 4,5,6 > 4 => 3
// i: arr = [11,13,17,23,29,31,7,5,2,3], k = 3, threshold = 5 => 6

func numOfSubarrays(arr []int, k int, threshold int) int {
	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}

	count := 0
	target := threshold * k
	if sum >= target {
		count++
	}

	for i := k; i < len(arr); i++ {
		sum = sum + arr[i] - arr[i-k]
		if sum >= target {
			count++
		}
	}
	return count
}
