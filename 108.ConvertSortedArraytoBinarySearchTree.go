/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sortedArrayToBST(nums []int) *TreeNode {
    return build(nums, 0, len(nums)-1)
}

 func build(nums []int, l, r int) *TreeNode {
        
    if l > r {return nil}
     
    m := (l+r)/2
    root := new(TreeNode)
    root.Val = nums[m]
    root.Left = build(nums, l, m-1)
    root.Right = build(nums, m+1, r)
    return root
} 