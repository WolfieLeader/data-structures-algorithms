package main

// TODO

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters) // [l, r]

	for left < right {
		mid := left + ((right - left) / 2)

		// Upper bound
		if target < letters[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return letters[left%len(letters)] // return first letter if not found
}
