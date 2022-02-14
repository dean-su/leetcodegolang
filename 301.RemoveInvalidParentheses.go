func removeInvalidParentheses(s string) []string {
    l, r := 0, 0
    res := []string{}
    for _, ch := range s{
        if ch == '(' { l++ }
        if l == 0 && ch == ')' {
            r++
        }else if ch == ')' {
            l--
        }
    }
    
    dfs(s, &res, 0, l, r)
    
    return res
}

func isValid(s string) bool {
    i := 0
    for _, ch := range s {
        if ch == '(' {i++}
        if ch == ')' {i--}
        if i<0 {return false}
    }
    return i == 0
}

func dfs(s string, res *[]string, start int, l, r int) {
    if l == 0 && r == 0 {
        if isValid(s) {*res = append(*res, s)}
        return
    }
    
    //for i, ch := range s {
    for i := start; i < len(s); i++ {
        if i != start && s[i] == s[i-1]  {continue}
        if s[i] == '(' || s[i] == ')' {
           
            curr := s
            curr = curr[:i] + curr[i+1:]
            if r>0 && s[i] == ')' {
                dfs(curr, res, i, l, r-1)
            }else if l > 0 && s[i] == '(' {
                dfs(curr, res, i, l-1, r)
            }
        }
    }
}