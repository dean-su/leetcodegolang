func numIslands(grid [][]byte) int {
    n, m , res := len(grid), len(grid[0]), 0
    
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == '1' {
                res ++
                dfs(grid, i, j)
            }
        }
    }
    
    return res
    
}

func dfs(grid [][]byte, x, y int) {
    if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0' {
        return 
    }
    
    grid[x][y] = '0'
    dfs(grid, x+1, y)
    dfs(grid, x-1, y)
    dfs(grid, x, y+1)
    dfs(grid, x, y-1)
}