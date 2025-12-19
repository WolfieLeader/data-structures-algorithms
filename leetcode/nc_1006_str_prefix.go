package main

import (
	"strconv"
	"strings"
)

type SolutionNC1006 struct{}

func (s *SolutionNC1006) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteByte('#')
		sb.WriteString(str)
	}
	return sb.String()
}

func (s *SolutionNC1006) Decode(encoded string) []string {
	out := make([]string, 0)
	i := 0

	for i < len(encoded) {
		j := i
		// find '#'
		for j < len(encoded) && encoded[j] != '#' {
			j++
		}
		if j == len(encoded) {
			return []string{} // bad
		}

		length, err := strconv.Atoi(encoded[i:j])
		if err != nil {
			return []string{} // bad
		}

		start := j + 1
		end := start + length
		if end > len(encoded) {
			return []string{} // bad
		}

		out = append(out, encoded[start:end])
		i = end
	}

	return out
}
