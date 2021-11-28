// https://programmers.co.kr/learn/courses/30/lessons/77484
package main

func main(lottos []int, win_nums []int) []int {
	lucky := []int{}
	result := []int{}
	count := 0
	max := 0
	min := 0

	for _, v := range lottos {
		if v == 0 {
			count++
			continue
		}
		if contains(win_nums, v) {
			lucky = append(lucky, v)
		}
	}

	if len(lucky) <= 1 {
		min = 6
	} else if len(lucky) == 2 {
		min = 5
	} else if len(lucky) == 3 {
		min = 4
	} else if len(lucky) == 4 {
		min = 3
	} else if len(lucky) == 5 {
		min = 2
	} else {
		min = 1
	}

	if len(lucky)+count == 6 {
		max = 1
	} else if len(lucky)+count == 5 {
		max = 2
	} else if len(lucky)+count == 4 {
		max = 3
	} else if len(lucky)+count == 3 {
		max = 4
	} else if len(lucky)+count == 2 {
		max = 5
	} else {
		max = 6
	}

	result = []int{max, min}
	return result
}

func contains(list []int, i int) bool {
	for _, value := range list {
		if value == i {
			return true
		}
	}
	return false
}
