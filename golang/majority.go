package main

func main() {
	// nums := []int{3, 2, 3}
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	majorityElement(nums)
}

func majorityElement(nums []int) int {
	numsMap := make(map[int]int)
	result := 0

	for _, v := range nums {
		numsMap[v]++
		if numsMap[v] > len(nums)/2 {
			result = v
			break
		}
	}
	return result
}
