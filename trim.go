package gosupport

import "strings"

func TrimStar(s string) string {
	s1 := strings.Trim(s, " *")
	return s1
}

func TrimLeftStar(s string) string {
	s1 := strings.TrimLeft(s, " *")
	return s1
}

func TrimRightStar(s string) string {
	s1 := strings.TrimRight(s, " *")
	return s1
}
