package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var input int
	var max int
	var index int
	for i := 0; i < 9; i++ {
		fmt.Fscanln(reader, &input)
		if max < input {
			max = input
			index = i
		}
	}
	fmt.Println(max)
	fmt.Println(index + 1)
}
