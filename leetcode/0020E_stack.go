package main

// Stack
//
// Use a stack to track opening brackets.
// When a closing bracket appears, it must match the most recent opening one (LIFO)
// If it doesn’t match or the stack is empty, the string is invalid.
//
// A map defines which opening bracket corresponds to each closing bracket.
// At the end, the stack must be empty for the string to be valid.
//
// Walkthrough: "{[()]}"
// push '{'
// push '['
// push '('
// ')' -> matches '(' -> pop
// ']' -> matches '[' -> pop
// '}' -> matches '{' -> pop
// stack empty -> valid

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		currChar := s[i]
		if currChar == '(' || currChar == '{' || currChar == '[' {
			stack = append(stack, currChar) // Push

		} else if openBracket, ok := pairs[currChar]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != openBracket { // Peek
				return false
			}

			stack = stack[:len(stack)-1] // Pop

		} else {
			return false // Not a bracket
		}
	}

	return len(stack) == 0
}
