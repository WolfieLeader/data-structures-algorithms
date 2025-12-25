package main

// TODO:

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		currFirstStrChar := strs[0][i]

		// NOTE: LOOP to see that all strs contain the same as the first
		for _, str := range strs[1:] {
			// Out of bounds or mismatch
			if len(str) <= i || str[i] != currFirstStrChar {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
