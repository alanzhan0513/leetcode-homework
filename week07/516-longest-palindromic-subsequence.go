func longestPalindromeSubseq(s string) int {
	length := len(s)
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}
	for l := 1; l <= length; l++ {
		for i := 0; i <= length - l; i++ {
			j := i + l - 1
			if i == j {
				dp[i][j] = 1
				continue
			}
			if s[i] == s[j] {
				dp[i][j] = dp[i + 1][j - 1] + 2
			} else {
				dp[i][j] = max(dp[i + 1][j], dp[i][j - 1])
			}
		}
	}
	return dp[0][length - 1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}