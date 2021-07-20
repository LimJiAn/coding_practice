// https://programmers.co.kr/learn/courses/30/lessons/81302

package main

import (
	"fmt"
)

func checkDistance(i int, j int, place []string) bool {
	if j+1 < 5 && place[i][j+1] == 'P' {
		return false
	}

	if i+1 < 5 && place[i+1][j] == 'P' {
		return false
	}

	if i+1 < 5 && j+1 < 5 && !(place[i+1][j] == 'X' && place[i][j+1] == 'X') && place[i+1][j+1] == 'P' {
		return false
	}

	if j+2 < 5 && place[i][j+1] != 'X' && place[i][j+2] == 'P' {
		return false
	}

	if i+2 < 5 && place[i+1][j] != 'X' && place[i+2][j] == 'P' {
		return false
	}

	if i-1 >= 0 && j+1 < 5 && place[i-1][j+1] == 'P' && !(place[i][j+1] == 'X' && place[i-1][j] == 'X') {
		return false
	}

	return true
}

func checkPlace(place []string) bool {
	for i, _ := range place {
		for j := 0; j < 5; j++ {
			if place[i][j] == 'P' && !checkDistance(i, j, place) {
				return false
			}
		}
	}
	return true
}

func solution(places [][]string) []int {
	var result []int
	for _, place := range places {
		if checkPlace(place) {
			result = append(result, 1)
		} else {
			result = append(result, 0)
		}
	}
	return result
}

func main() {
	// places 원소의 길이(대기실 가로 길이) = 5
	// P는 응시자가 앉아있는 자리를 의미합니다.
	// O는 빈 테이블을 의미합니다.
	// X는 파티션을 의미합니다.

	// var P string
	// var O string
	// var X string

	places := [][]string{{"POOOP", "OXXOX", "OPXPX", "OOXOX", "POXXP"}, {"POOPX", "OXPXP", "PXXXO", "OXXXO", "OOOPP"}, {"PXOPX", "OXOXP", "OXPOX", "OXXOP", "PXPOX"}, {"OOOXX", "XOOOX", "OOOXX", "OXOOX", "OOOOO"}, {"PXPXP", "XPXPX", "PXPXP", "XPXPX", "PXPXP"}}
	fmt.Println(solution(places))
}
