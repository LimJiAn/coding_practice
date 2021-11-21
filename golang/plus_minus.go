// https://programmers.co.kr/learn/courses/30/lessons/76501?language=go
package main

func main() {
	solution([]int{4, 7, 12}, []bool{true, false, true})
}

func solution(absolutes []int, signs []bool) int {
	var (
		result []int
		sum    int
	)
	for i, v := range absolutes {
		if !signs[i] {
			v = -v
		}
		result = append(result, v)
	}

	for _, v := range result {
		sum += v
	}
	return sum
}
