// https://programmers.co.kr/learn/courses/30/lessons/86051

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 6, 7, 8, 0}
	total := 0
	sum := 0

	for i := 0; i < 10; i++ {
		total += i
	}
	for _, v := range numbers {
		if contains(v) {
			sum += v
		}
	}
	fmt.Println(total - sum)
}

func contains(n int) bool {
	for i := 0; i < 10; i++ {
		if i == n {
			return true
		}
	}
	return false
}
