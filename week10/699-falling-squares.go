func fallingSquares(positions [][]int) []int {
    height := make([]int, len(positions))
	for i, v := range positions {
		l1, r1, curHeight := v[0], v[0]+v[1]-1, v[1]
		height[i] = curHeight
		for j, vv := range positions[:i] {
			l2, r2 := vv[0], vv[0]+vv[1]-1
			if !(r1 < l2 || l1 > r2) {
				height[i] = max(height[i], height[j]+curHeight)
			}
		}
	}
	for i := 1; i < len(height); i++ {
		height[i] = max(height[i], height[i-1])
	}
	return height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
