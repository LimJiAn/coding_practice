package main

import (
	"strings"
)

func main() {
	text := "leet code"
	brokenLetters := "lt"
	canBeTypedWords(text, brokenLetters)
}

func canBeTypedWords(text string, brokenLetters string) int {
	textArray := strings.Split(text, " ")
	brokenLettersArray := strings.Split(brokenLetters, "")
	typedWords := len(textArray)
	for _, v := range textArray {
		sign := false
		for _, s := range brokenLettersArray {
			if strings.Contains(v, s) {
				if sign == true {
					continue
				}
				typedWords -= 1
				sign = true
			}
		}
	}
	return typedWords
}
