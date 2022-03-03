
func uniquePaths(m int, n int) int {
    grid := make([][]int, m)
    for i := 0; i < m; i++ {
        grid[i] = make([]int, n)
        for j := 0; j < n; j++ {
            if i == 0 || j == 0{
                grid[i][j] = 1
                continue
            }
            
            grid[i][j] = grid[i-1][j] + grid[i][j-1]
        }
    }
    
    return grid[m-1][n-1]
}