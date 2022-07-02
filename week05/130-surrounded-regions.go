func solve(board [][]byte) {
    m, n := len(board), len(board[0])
    var dfs func(r, c int)
    dfs = func(r, c int) {
        if r < 0 || r >= m || c < 0 || c >= n { return }
        if board[r][c] != 'O' { return }
        board[r][c] = '$'
        dfs(r+1, c)
        dfs(r-1, c)
        dfs(r, c+1)
        dfs(r, c-1)
    }
    
    for r := 0; r < m; r++ {
        dfs(r, 0)
        dfs(r, n-1)
    }
    for c := 0; c < n; c++ {
        dfs(0, c)
        dfs(m-1, c)
    }
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if board[r][c] != '$' {
                board[r][c] = 'X'
            } else {
                board[r][c] = 'O'
            }
        }
    }
}