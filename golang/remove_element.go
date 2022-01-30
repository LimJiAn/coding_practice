// https://leetcode.com/problems/remove-element/

package main

func main() {
	// nums := []int{3, 2, 2, 3}
	// val := 3
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	removeElement(nums, val)
}

func removeElement(nums []int, val int) int {
	i := 0
	n := len(nums)

	for i < n {
		if nums[i] == val {
			n--
			nums[i] = nums[n]
		} else {
			i++
		}
	}

	return n
}
