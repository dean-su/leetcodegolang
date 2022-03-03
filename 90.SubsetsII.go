/*
90. Subsets II
Medium

Given an integer array nums that may contain duplicates, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

 

Example 1:

Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
Example 2:

Input: nums = [0]
Output: [[],[0]]
 

Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10
*/
func subsetsWithDup(nums []int) [][]int {
    curr := []int{}
    res := [][]int{}
    sort.Ints(nums)
    dfs(nums, curr, &res, 0)
    return res
}

func dfs(nums, curr []int, res *[][]int, s int) {
    *res = append(*res, append([]int{}, curr...))
    
    for i := s; i < len(nums); i++ {
        if i > s && nums[i] == nums[i-1] {continue}
        curr = append(curr, nums[i])
        dfs(nums, curr, res, i+1)
        curr = curr[:len(curr)-1]
    }
}