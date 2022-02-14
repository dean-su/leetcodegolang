/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced(root *TreeNode) bool {
    if root == nil {return true}
    
    return abs(height(root.Left) - height(root.Right)) < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
    if root == nil {return 0}
    
    return 1 + max(height(root.Left), height(root.Right))
}

func abs(a int) int {
    if a < 0 {
        return -a
    }else {
        return a
    }
}

func max(a, b int) int {
    if a < b {
        return b
    }else {
        return a
    }
}