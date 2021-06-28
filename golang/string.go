// https://www.acmicpc.net/problem/9086

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var length int
	var str string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &length)
	for i := 0; i < length; i++ {
		fmt.Fscanln(reader, &str)
		fmt.Println(string(str[0]) + string(str[len(str)-1]))
	}
}
