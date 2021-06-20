// https://www.acmicpc.net/problem/1330

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &a, &b)

	if a > b {
		fmt.Println(">")
	} else if a < b {
		fmt.Println("<")
	} else if a == b {
		fmt.Print("==")
	}
}
