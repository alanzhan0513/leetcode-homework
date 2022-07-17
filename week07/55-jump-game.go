func canJump(nums []int) bool {
	aim := len(nums) - 1
	for i:= len(nums) - 2; i >= 0; i-- {
		maxJump := min(i + nums[i], len(nums) - 1)
		for j := i + 1; j <= maxJump; j++ {
			if j == aim{
				aim = i
			}
		}
	}
	if aim == 0 {
		return true
	}
	return false
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}