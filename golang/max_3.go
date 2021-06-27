// https://www.acmicpc.net/problem/10818

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var length, input int
	var list []int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &length)

	for i := 0; i < length; i++ {
		fmt.Fscanf(reader, "%d", &input)
		//fmt.Fscanln(reader, &input) 비교
		list = append(list, input)
	}
	sort.Ints(list)
	fmt.Println(list[0], list[length-1])
}
