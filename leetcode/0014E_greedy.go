package main

// TODO:

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		currFirstStrChar := strs[0][i]

		// NOTE: LOOP to see that all strs contain the same as the first
		for j := 1; j < len(strs); j++ {
			// Out of bounds or mismatch
			if len(strs[j]) <= i || strs[j][i] != currFirstStrChar {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
