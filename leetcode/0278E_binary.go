package main

func firstBadVersion(n int) int {
	left, right := 1, n // [l, r)
	for left < right {
		mid := left + ((right - left) / 2)

		// Lower bound
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func isBadVersion(version int) bool

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
