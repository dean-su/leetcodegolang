func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    
    d := make(map[byte][]string,10)
    d[0] = []string{""}
    d[1] = append(d[1], )
    d[2] = append(d[2], "a", "b", "c")
    d[3] = append(d[3], "d", "e", "f")
    d[4] = append(d[4], "g", "h", "i")
    d[5] = append(d[5], "j", "k", "l")
    d[6] = append(d[6], "m", "n", "o")
    d[7] = append(d[7], "p", "q", "r", "s")
    d[8] = append(d[8], "t", "u", "v")
    d[9] = append(d[9], "w", "x", "y", "z")
    
    res := []string{}
    var curr string
    dfs(digits, d, 0, curr, &res)
    
    return res
}

func dfs(digits string, d  map[byte][]string, l int, curr string, res *[]string) {
    if l == len(digits) {
        *res = append(*res, curr)
        return
    }
    
    for _, c := range d[digits[l]-'0'] {
        curr += c
        dfs(digits, d, l+1, curr, res)
        curr = curr[:len(curr)-1]
    }
}