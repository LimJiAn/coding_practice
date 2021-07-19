// https://www.acmicpc.net/problem/11399

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var length, input, sum int
	var list []int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &length)
	for i := 0; i < length; i++ {
		fmt.Fscan(reader, &input)
		list = append(list, input)
	}
	sort.Ints(list)
	for i := 1; i < length; i++ {
		list[i] = list[i-1] + list[i]
		sum += list[i]
	}
	// list[1] = list[0] + list[1]
	// list[2] = list[1] + list[2]
	//		...
	// 첫째항 안 더해줌
	fmt.Println(sum + list[0])
}
