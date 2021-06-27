// https://www.acmicpc.net/problem/3052

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var input int
	var remainder_list []int
	for i := 0; i < 10; i++ {
		fmt.Fscanln(reader, &input)
		remainder := input % 42
		if !contains(remainder_list, remainder) {
			remainder_list = append(remainder_list, remainder)
		}
	}
	fmt.Println(len(remainder_list))
}

func contains(list []int, i int) bool {
	for _, value := range list {
		if value == i {
			return true
		}
	}
	return false
}
