package main

type BracketType uint8

const (
	NotABracket    BracketType = 0
	OpeningBracket BracketType = 1
	ClosingBracket BracketType = 2
)

func bracketType(b byte) BracketType {
	switch b {
	case '(', '{', '[':
		return OpeningBracket
	case ')', '}', ']':
		return ClosingBracket
	default:
		return NotABracket
	}
}

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		currChar := s[i]
		typ := bracketType(currChar)
		if typ == NotABracket {
			return false
		}
		if typ == OpeningBracket {
			stack = append(stack, currChar)
			continue
		}
		// Closing Bracket
		n := len(stack)
		if n == 0 {
			return false
		}

		top := stack[n-1]
		stack = stack[:n-1]

		if (top == '(' && currChar != ')') || (top == '{' && currChar != '}') || (top == '[' && currChar != ']') {
			return false
		}
	}
	return len(stack) == 0
}

// ----------------------------------------------------------------------------------

func optimizedIsValid(s string) bool {
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
