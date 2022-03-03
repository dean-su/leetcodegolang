func combinationSum2(candidates []int, target int) [][]int {
    var curr []int
    var res [][]int
    sort.Ints(candidates)
    dfs(candidates, target, 0, curr, &res)
    return res
}

func dfs(cans []int, t, s int, curr []int, res *[][]int) {
    if t == 0 {
        temp := make([]int, len(curr))
        copy(temp, curr)
        *res = append(*res, temp)
        return
    }
    
    for i := s; i < len(cans); i++ {
        if cans[i] > t { return }
        if i > s && cans[i] == cans[i-1] { continue }
        curr = append(curr, cans[i])
        dfs(cans, t - cans[i], i+1, curr, res)
        curr = curr[:len(curr)-1]
    }
}