package main

// TODO:

// XOR (Exclusive OR) properties:
// x ^ x = 0   (duplicates cancel out)
// x ^ 0 = x   (zero does nothing)
// XOR is associative and commutative
//
// By XORing all numbers, every duplicated value cancels itself,
// leaving only the number that appears once.

func singleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
	}
	return result
}

// Walk through input: [3,1,5,1,3]
// 0 ^ 3 = 3
// 3 ^ 1 = 2
// 2 ^ 5 = 7
// 7 ^ 1 = 6
// 6 ^ 3 = 5
