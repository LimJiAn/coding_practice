// https://leetcode.com/problems/adding-spaces-to-a-string/

package main

import (
	"bytes"
)

func main() {
	// s := "LeetcodeHelpsMeLearn"
	// spaces := []int{8, 13, 15}
	// s := "icodeinpython"
	// spaces := []int{1, 5, 7, 9}
	// s := "p"
	// spaces := []int{0}
	s := "ezixkFLjdbxrDowLVGYvdtBrguAAJVM"
	spaces := []int{23}

	addSpaces(s, spaces)
}

// func addSpaces(s string, spaces []int) string {
// 	if len(spaces) == 1 {
// 		a := s[0:spaces[0]]
// 		b := s[spaces[0]:]
// 		return a + " " + b
// 	}

// 	var slice []string
// 	for i := 1; i < len(spaces); i++ {
// 		a := s[spaces[i-1]:spaces[i]]
// 		slice = append(slice, a)
// 	}
// 	for i := 0; i < len(slice); i++ {
// 		tmp := strings.Split(s, slice[i])
// 		fmt.Println(tmp)
// 		if i == 0 {
// 			slice = append([]string{tmp[0]}, slice...)
// 		} else if i == len(slice)-1 {
// 			slice = append(slice, tmp[len(tmp)-1])
// 			break
// 		}
// 	}
// 	str := strings.Join(slice, " ")
// 	return str
// }

func addSpaces(s string, spaces []int) string {
	var ans bytes.Buffer
	l := 0

	for i := 0; i < len(spaces); i++ {
		ans.WriteString(s[l:spaces[i]] + " ")
		l = spaces[i]
	}

	// ans.WriteString(s[l:len(s)])
	ans.WriteString(s[l:])
	// fmt.Println(ans.String())

	return ans.String()
}
