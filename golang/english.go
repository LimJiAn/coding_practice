package main

import (
	"strconv"
	"strings"
)

func main(s string) int {
	table := []string{"zero", "one", "two", "three",
		"four", "five", "six", "seven", "eight", "nine"}

	for i, v := range table {
		s = strings.Replace(s, v, strconv.Itoa(i), -1)
	}
	result, _ := strconv.Atoi(s)
	return result
}
