package main

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pairs := map[byte]byte{')': '(', '}': '{', ']': '['}
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		currChar := s[i]

		// Closing
		if open, ok := pairs[currChar]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}

		// Opening
		if currChar == '(' || currChar == '{' || currChar == '[' {
			stack = append(stack, currChar)
		} else {
			return false // Something else
		}
	}

	return len(stack) == 0
}
