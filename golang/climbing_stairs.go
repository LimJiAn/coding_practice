package main

func main() {
	// n := 3
	// n := 4
	// n := 5
	n := 6
	// n := 7
	climbStairs(n)
}

func climbStairs(n int) int {
	a, b := 1, 1
	for i := 1; i < n; i++ {
		next := a + b
		a, b = b, next
	}
	return b
}
