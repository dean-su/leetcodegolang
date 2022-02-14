func generateParenthesis(n int) []string {
    res := make([]string, 0)
    var curr string
    dfs(&res, curr, n, n)
    return res
}

func dfs(res *[]string, curr string, l, r int) {
    if l == 0 && r == 0 {
        *res = append(*res, curr)
        return
    }
    
    if l > 0 {
        dfs(res, curr+"(", l-1, r)
    }
    
    if r >l {
        dfs(res, curr+")", l, r-1)
    }
}