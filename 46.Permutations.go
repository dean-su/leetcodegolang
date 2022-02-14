func permute(nums []int) [][]int {
    res := [][]int{}
    permutation(nums, &res, 0)
    
    return res
}

func permutation(nums []int, res *[][]int, d int) {
    if d == len(nums) {
        dummy := make([]int, len(nums))
        copy(dummy, nums)
        *res = append(*res, dummy)
        return
    }
    
    for i := d; i < len(nums); i++ {
        nums[i], nums[d] = nums[d], nums[i]
        permutation(nums, res, d+1)
        nums[i], nums[d] = nums[d], nums[i]
    }
}