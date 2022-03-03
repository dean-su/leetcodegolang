func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if obstacleGrid[0][0] == 1 {return 0}
    m := len(obstacleGrid) 
    n := len(obstacleGrid[0])
    f := make([][]int, m)

    for i := 0; i < m; i++ {
        f[i] = make([]int, n)
        f[0][0] = 1
        l, r := 0, 0
        for j := 0; j < n; j++ {
            if obstacleGrid[i][j] == 1 || (i==0 && j==0) {
                continue
            }else {
                if i > 0 {
                    l = f[i-1][j]
                }
                if j > 0 {
                    r = f[i][j-1]
                }
                f[i][j] = l+r
            }
        }
    }
    return f[m-1][n-1]
}