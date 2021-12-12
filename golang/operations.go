// https://leetcode.com/problems/final-value-of-variable-after-performing-operations/

package main

import (
	"strings"
)

func main() {
	operations := []string{"--X", "X++", "X++"}
	finalValueAfterOperations(operations)
}

func finalValueAfterOperations(operations []string) int {
	sum := 0
	if len(operations) == 0 {
		return sum
	}
	for _, v := range operations {
		if strings.Contains(v, "+") {
			sum += 1
		} else {
			sum -= 1
		}
	}
	return sum
}
