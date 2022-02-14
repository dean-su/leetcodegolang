/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isValidBST(root *TreeNode) bool {
    return inOrder(root, nil, nil)
}

func inOrder(root, min, max *TreeNode) bool {
    if root == nil {return true}
    
    if min != nil && root.Val <= min.Val {
        return false
    }
    
    if max != nil && root.Val >= max.Val {
        return false
    }
    
    return inOrder(root.Left, min, root) && inOrder(root.Right, root, max) 
}