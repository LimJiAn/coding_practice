// https://www.acmicpc.net/problem/2789

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	str := "CAMBRIDGE"
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &input)
	for _, i := range str {
		input = strings.Replace(input, string(i), "", -1) // -1 전체 (양수 반복)
	}
	fmt.Println(input)
}
