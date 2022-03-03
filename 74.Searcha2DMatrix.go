func searchMatrix(matrix [][]int, target int) bool {
    l := 0;
    r := len(matrix) * len(matrix[0])
    c := len(matrix[0])
    
    for l < r {
        m := (l+r)/2
        if matrix[m/c][m%c] == target {
            return true
        }else if matrix[m/c][m%c] > target {
            r = m
        }else {
            l = m+1
        }
    }
    return false
}