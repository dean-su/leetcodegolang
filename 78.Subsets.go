/*
78. Subsets
Medium

Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

 

Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
Example 2:

Input: nums = [0]
Output: [[],[0]]
 

Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10
All the numbers of nums are unique.
*/
func subsets(nums []int) [][]int {
    curr := []int{}
    res :=[][]int{}
    for i := 0; i <= len(nums); i++ {
        dfs(nums, curr, &res, i, 0)
    }
    return res
}

func dfs(nums, curr []int, res *[][]int, n, s int) {
    if len(curr) == n {
        *res = append(*res, append([]int{}, curr...))
        return
    }
    
    for i := s; i < len(nums); i++ {
        curr = append(curr, nums[i])
        dfs(nums, curr, res, n, i+1)
        curr = curr[:len(curr)-1]
    }
    
}