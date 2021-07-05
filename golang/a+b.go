// https://www.acmicpc.net/problem/10950

package main

import "fmt"

func main() {
	var input, a, b int
	fmt.Scan(&input)
	for i := 0; i < input; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
