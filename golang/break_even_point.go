// https://www.acmicpc.net/problem/1712

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c, n int
	// a 고정 비용
	// b 가변 비용
	// c 판매 비용
	// n 갯수
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &a, &b, &c)

	// a+(b*n) == (c * n)	양 변을 n 으로 나눔
	// a/n + b == c
	// a/n = c - b
	// n = a/(c-b)
	if (c - b) <= 0 {
		fmt.Println(-1)
	} else {
		n = a / (c - b)
		fmt.Print(n + 1)
	}
}
