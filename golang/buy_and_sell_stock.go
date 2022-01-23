package main

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}
	// prices = []int{7, 6, 4, 3, 1}
	maxProfit(prices)
}

func maxProfit(prices []int) int {
	min := prices[0]
	result := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			tmp := prices[i] - min
			if tmp > result {
				result = tmp
			}
		}
	}
	return result
}
