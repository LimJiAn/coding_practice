package main

// https://leetcode.com/problems/longest-common-prefix/

import (
	"strings"
)

func main() {
	// strs := []string{"flower", "flow", "flight"}
	strs := []string{"dog", "racecar", "car"}
	longestCommonPrefix(strs)
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for _, str := range strs {
		if len(prefix) > len(str) {
			prefix = str
		}
	}
	for true {
		for i, str := range strs {
			if !strings.HasPrefix(str, prefix) {
				prefix = prefix[:len(prefix)-1]
				break
			}
			if i == len(strs)-1 {
				return prefix
			}
		}
	}
	return prefix
}
