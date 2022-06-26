func searchMatrix(matrix [][]int, target int) bool {
    y := len(matrix)
    if y == 0 {
        return false
    }
    x := len(matrix[0])
    
    left := 0
    right := x * y - 1
    for right >= left {
        mid := (left + right) / 2
        i := mid / x
        j := mid % x
        element := matrix[i][j]
        if element == target {
            return true
        }
        if element > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return false
}