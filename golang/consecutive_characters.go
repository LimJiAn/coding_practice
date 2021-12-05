// https://leetcode.com/problems/consecutive-characters/

package main

func main() {
	s := "j"
	// s = "abbcccddddeeeeedcba"
	// s = "triplepillooooow"
	// s = "hooraaaaaaaaaaay"
	maxPower(s)
}

func maxPower(s string) int {
	result, count := 1, 1

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			count++
		} else {
			count = 1
		}
		if result < count {
			result = count
		}
	}

	return result
}
