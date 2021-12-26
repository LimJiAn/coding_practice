package main

func main() {
	nums := []int{1, 2, 3, 1}
	// nums := []int{1, 2, 3, 4}
	containsDuplicate(nums)
}

func containsDuplicate(nums []int) bool {
	intMap := make(map[int]bool)
	check := []int{}
	for _, num := range nums {
		if _, val := intMap[num]; !val {
			intMap[num] = true
			check = append(check, num)
		} else {
			intMap[num] = false
		}
	}
	if len(nums) == len(check) {
		return false
	} else {
		return true
	}
}
