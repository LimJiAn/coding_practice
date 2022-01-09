package main

// https://leetcode.com/problems/maximum-subarray/

func main() {
	// nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	nums := []int{5, 4, -1, 7, 8}
	// nums := []int{1}
	maxSubArray(nums)
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := nums[0]
	result := 0
	check := 0
	for check < len(nums) {
		result += nums[check]
		if result > maxSum {
			maxSum = result
		}
		if result < 0 {
			result = 0
		}
		check++
	}
	return maxSum
}
