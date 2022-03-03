/*
Medium

Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

 

Example 1:


Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.
Example 2:

Input: grid = [[1,2,3],[4,5,6]]
Output: 12
*/

func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 {continue}
            if i == 0 {
                grid[i][j] += grid[i][j-1]
            } else if j == 0 {
                grid[i][j] += grid[i-1][j]
            } else {
                grid[i][j] += min(grid[i][j-1], grid[i-1][j])
            }
        }
    }
    return grid[m-1][n-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }else {
        return b
    }
}