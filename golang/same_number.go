// https://programmers.co.kr/learn/courses/30/lessons/12906

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 5, 5, 1, 2, 3, 4, 4, 4, 4}
	new_arr := []int{}
	for _, v := range arr {
		if !contains(new_arr, v) {
			new_arr = append(new_arr, v)
		}
	}
	fmt.Print(new_arr)
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
