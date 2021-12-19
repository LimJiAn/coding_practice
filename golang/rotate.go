package main

import "strings"

func main() {
	s := "abcde"
	goal := "cdeab"
	rotateString(s, goal)
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	} else {
		s = s + s
		// abcdeabcde
		if strings.Contains(s, goal) {
			return true
		} else {
			return false
		}
	}
}
