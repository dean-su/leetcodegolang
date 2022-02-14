/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func recoverTree(root *TreeNode)  {
    var pre, first, second *TreeNode
    var inorder func(*TreeNode)
    inorder = func(root *TreeNode) {
        if root == nil {return}
        inorder(root.Left)
        if pre != nil && pre.Val > root.Val {
            if first == nil {first = pre}
            second = root
        }
        pre = root
        inorder(root.Right)
    }
    inorder(root)
    first.Val, second.Val = second.Val, first.Val
}