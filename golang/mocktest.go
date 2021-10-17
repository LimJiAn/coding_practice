// https://programmers.co.kr/learn/courses/30/lessons/42746?language=go

package main

func solution(answers []int) []int {
	people1 := [5]int{1, 2, 3, 4, 5}
	people2 := [8]int{2, 1, 2, 3, 2, 4, 2, 5}
	people3 := [10]int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}

	var correct1, correct2, correct3 int = 0, 0, 0
	for i, v := range answers {
		if v == people1[i%5] {
			correct1++
		}
		if v == people2[i%8] {
			correct2++
		}
		if v == people3[i%10] {
			correct3++
		}
	}

	var max_correct_people int
	max_correct_people = max(correct1, correct2)
	max_correct_people = max(max_correct_people, correct3)

	who_max_correct_peple := []int{}
	if max_correct_people == correct1 {
		who_max_correct_peple = append(who_max_correct_peple, 1)
	}

	if max_correct_people == correct2 {
		who_max_correct_peple = append(who_max_correct_peple, 2)
	}

	if max_correct_people == correct3 {
		who_max_correct_peple = append(who_max_correct_peple, 3)
	}

	return who_max_correct_peple
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
