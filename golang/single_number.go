package main

func main() {
	nums := []int{2, 2, 1}
	// nums := []int{4,1,2,1,2}
	// nums := []int{1}
	singleNumber(nums)
}

func singleNumber(nums []int) int {
	result := 0
	numsMap := make(map[int]int)

	for _, v := range nums {
		numsMap[v]++
	}

	for k, v := range numsMap {
		if v == 1 {
			result = k
			break
		}
	}
	return result
}
