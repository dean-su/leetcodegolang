func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    path := make([]int, 0)
    used := make(map[int]bool, len(nums))
    
    dfs(nums, &res, path, used, 0)
    
    return res
}

func dfs(nums []int, res *[][]int, path []int, used map[int]bool, d int) {
    if d == len(nums) {
        dummy := make([]int, len(path))
        copy(dummy, path)
        *res = append(*res, dummy)
        return
    }
    
    for i := 0; i < len(nums); i++ {
        if used[i] == true {continue}
        if i > 0 && nums[i] == nums[i-1] && !used[i-1] {continue}
        
        used[i] = true
        path = append(path, nums[i])
        dfs(nums, res, path, used, d+1)
        used[i] = false
        path = path[:len(path)-1]
    }
}