// https://www.acmicpc.net/problem/9498

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &input)
	if input >= 90 && input <= 100 {
		fmt.Println("A")
	} else if input >= 80 && input <= 89 {
		fmt.Println("B")
	} else if input >= 70 && input <= 79 {
		fmt.Println("C")
	} else if input >= 60 && input <= 69 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
