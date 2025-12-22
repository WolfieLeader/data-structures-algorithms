package main

// TODO:

func isSubsequence(s, t string) bool {
	p1, p2 := 0, 0
	for p2 < len(t) && p1 < len(s) {
		if s[p1] == t[p2] {
			p1++
		}
		p2++
	}
	return p1 == len(s)
}
