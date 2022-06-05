func maximalRectangle(matrix [][]byte) int {
    matrixHeight := len(matrix)
    if matrixHeight == 0 {
        return 0
    }
    
    matrixWidth := len(matrix[0])
    if matrixWidth == 0 {
        return 0
    }
    
    dp := make([][]int, matrixHeight)
	for i := 0; i < matrixHeight; i++ {
		dp[i] = make([]int, matrixWidth)
	}

	for j := 0; j < matrixWidth; j++ {
		dp[0][j] = int(matrix[0][j] - '0')
		for i := 1; i < matrixHeight; i++ {
			if matrix[i][j] == '1' {
				dp[i][j] = dp[i-1][j] + 1
			}
		}
	}

	max := 0
	for i := 0; i < matrixHeight; i++ {
		tmp := largestRectangleArea(dp[i])
		if max < tmp {
			max = tmp
		}
	}

	return max
}

type rect struct {
    width   int
    height  int
}

func largestRectangleArea(heights []int) int {
    stack := []rect{}
    max := 0
    heights = append(heights, 0)
    for _, height := range heights {
        accumulatedWidth := 0
        
        for len(stack) != 0 && stack[len(stack) - 1].height >= height {
            top := stack[len(stack) - 1]
            accumulatedWidth += top.width
            if accumulatedWidth * top.height > max {
                max = accumulatedWidth * top.height
            }
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, rect{accumulatedWidth + 1, height})
    }
    
    return max
}