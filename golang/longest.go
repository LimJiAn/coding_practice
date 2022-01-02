package main

func main() {
	// s := "abcabcbb"
	s := "pwwkew"
	// s := "abcdefgasd"
	lengthOfLongestSubstring(s)
}

func lengthOfLongestSubstring(s string) int {
	var result int
	slice := []string{}

	if len(s) == 0 {
		return result
	} else {
		for _, v := range s {
			for j, vv := range slice {
				if vv == string(v) {
					slice = slice[j+1:]
					break
				}
			}
			slice = append(slice, string(v))
			if len(slice) > result {
				result = len(slice)
			}
		}
	}
	return result
}
