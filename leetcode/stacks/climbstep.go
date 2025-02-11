package stacks

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	step1 := 1
	step2 := 2

	for i := 3; i <= n; i++ {
		curr := step1 + step2
		step1 = step2
		step2 = curr
	}
	return n

}
