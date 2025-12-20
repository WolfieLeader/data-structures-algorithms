package main

// I: nums = [1,7,3,6,5,6] => Ex: 1+7+3 == 5+6 => O: 3
// I: nums = [1,2,3] => O: -1
// I: nums = [2,1,-1] => O: 0

// total = leftSum + nums[i] + rightSum
//                 V
// rightSum = total - leftSum - nums[i]

func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	leftSum := 0
	for i, v := range nums {
		rightSum := total - leftSum - v

		if leftSum == rightSum {
			return i
		}

		leftSum += v
	}

	return -1
}
