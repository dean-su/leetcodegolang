/*Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

 

Example 1:


Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
*/

func exist(board [][]byte, word string) bool {
    if len(board) == 0 {return false}
    w, h := len(board[0]), len(board)
    for i := 0; i < w; i++ {
        for j := 0; j < h; j++ {
            if search(board, word, 0, i, j) == true {return true}
        }
    }
    return false
}

func search(board [][]byte, word string, d, x, y int) bool {
    if x < 0 || x == len(board[0]) || y < 0 || y == len(board) || word[d] != board[y][x] {
        return false
    }
    if d == len(word)-1 {return true}
    curr := board[y][x]
    board[y][x] = 0
    found := search(board, word, d+1, x+1, y) || search(board, word, d+1, x-1, y) || search(board, word, d+1, x, y+1) || search(board, word, d+1, x, y-1)
    
    board[y][x] = curr
    return found
}