// https://programmers.co.kr/learn/courses/30/lessons/12948?language=go

package main

import "fmt"

func main() {
	number := "029123091"
	result := []rune{}
	star := []rune("*")
	for i, v := range number {
		if i >= len(number)-4 {
			result = append(result, v)
		} else {
			result = append(result, star...)
		}
	}
	fmt.Println(string(result))
}

// 제출 풀이
// func solution(phone_number string) string {
//     result := []rune{}
//     star := []rune("*")
//     for i, v := range phone_number {
//         if i >= len(phone_number)-4 {
//             result = append(result, v)
//         } else {
//             result = append(result, star...)
//         }
//     }
//     return string(result)
// }
