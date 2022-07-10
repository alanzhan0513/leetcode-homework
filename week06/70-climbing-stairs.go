func climbStairs(n int) int {
    if (n == 0 || n == 1) {
        return n
    }
    res := make([]int, n + 1)
    res[1] = 1
    res[2] = 2
    for i := 3; i <= n; i++ {
        res[i] = res[i - 1] + res[i - 2]
    }
    return res[n]
}
