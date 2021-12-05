// https://leetcode.com/problems/two-sum/

package main

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	var m map[int]int
	m = map[int]int{}
	for i := 0; i < len(nums); i++ {
		one := target - nums[i]
		if _, ok := m[one]; ok {
			return []int{m[one], i}
		}
		m[nums[i]] = i
	}
	return nil
}
