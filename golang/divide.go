// https://programmers.co.kr/learn/courses/30/lessons/87389

package main

import "fmt"

func main() {
	var a = 10
	var i int

	for i = 1; i < a; i++ {
		if a%i == 1 {
			break
		}
	}
	fmt.Println(i)
}

// 제출 코드
// func solution(n int) int {
// 	var i int
// 	for i = 1; i < n; i++ {
// 		if n%i == 1 {
// 			break
// 		}
// 	}
// 	return i
// }
