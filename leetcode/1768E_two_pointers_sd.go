package main

import "strings"

func mergeAlternately(word1, word2 string) string {
	var sb strings.Builder
	sb.Grow(len(word1) + len(word2))

	i := 0
	for i < len(word1) || i < len(word2) {
		if i < len(word1) {
			sb.WriteByte(word1[i])
		}
		if i < len(word2) {
			sb.WriteByte(word2[i])
		}
		i++
	}

	return sb.String()
}
