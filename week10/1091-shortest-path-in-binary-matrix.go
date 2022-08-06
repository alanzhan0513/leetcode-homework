func shortestPathBinaryMatrix(grid [][]int) int {
    if grid[0][0] == 1 {
        return -1
    }
    
    result := 0
    n := len(grid)
    
    queue := make([][2]int, 0)
    queue = append(queue, [2]int{0, 0})
    
    visited := make(map[string]bool)
    visited[fmt.Sprintf("%v%v", 0, 0)] = true
    
    dirs := [8][2]int{{1,0},{1,1},{1,-1},{0,1},{0,-1},{-1,0},{-1,1},{-1,-1}}
    
    for len(queue) > 0 {
        size := len(queue)
        result++
        
        for i := 0; i < size; i++ {
            pop := queue[0]
            queue = queue[1:]
            
            // reaching the end
            if pop[0] == n-1 && pop[1] == n-1 {
                return result
            }
            
            for _, dir := range dirs {
                x := pop[0] + dir[0]
                y := pop[1] + dir[1]
                
                if x >= 0 && x < n && y >= 0 && y < n &&  grid[x][y] == 0 && !visited[fmt.Sprintf("%v%v", x, y)] {
                    visited[fmt.Sprintf("%v%v", x, y)] = true
                    queue = append(queue, [2]int{x, y})
                }
            }
        }
    }
    
    return -1
}
