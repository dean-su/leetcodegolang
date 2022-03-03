/*
113. Path Sum II
Medium

Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node values in the path equals targetSum. Each path should be returned as a list of the node values, not node references.

A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.

 

Example 1:


Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: [[5,4,11,2],[5,8,4,5]]
Explanation: There are two paths whose sum equals targetSum:
5 + 4 + 11 + 2 = 22
5 + 8 + 4 + 5 = 22
Example 2:


Input: root = [1,2,3], targetSum = 5
Output: []
Example 3:

Input: root = [1,2], targetSum = 0
Output: []
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, targetSum int) [][]int {
    curr := []int{}
    res := [][]int{}
    helper(root, targetSum, curr, &res)
    return res
}

func helper(root *TreeNode, targetSum int, curr []int, res *[][]int) {
    if root == nil {return}

    if root.Left == nil && root.Right == nil {
        if root.Val == targetSum {
			// temp := make([]int, len(curr))
            // copy(temp, curr)
            // *res = append(*res, temp)
            curr = append(curr, root.Val)
            *res = append(*res, append([]int{}, curr...))
            curr = curr[:len(curr)-1]  
        } 

        return
    }
    
    curr = append(curr, root.Val)
    newSum := targetSum - root.Val
    helper(root.Left, newSum, curr, res)
    helper(root.Right, newSum, curr, res)
    curr = curr[:len(curr)-1]
}