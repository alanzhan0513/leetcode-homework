func numSquares(n int) int {
    if n < 0 {
		return 0
	}
	if n <= 1 {
		return n
	}
	dp := make([]int, n + 1)

	for i := 1; i < n + 1; i++ {
        dp[i] = i
		for j := 1; i - j * j >= 0; j++ {
			dp[i] = min(dp[i], dp[i - j * j] + 1)
		}
	}

	return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
