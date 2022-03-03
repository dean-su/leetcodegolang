/*
easy
Given the root of a binary tree, return all root-to-leaf paths in any order.

A leaf is a node with no children.

 

Example 1:


Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]
Example 2:

Input: root = [1]
Output: ["1"]
 

Constraints:

The number of nodes in the tree is in the range [1, 100].
-100 <= Node.val <= 100
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func binaryTreePaths(root *TreeNode) []string {
    res := []string{}
    helper(root, "", &res)
    return res
}

func helper(root *TreeNode, curr string, res *[]string) {
    if root == nil {return}
    curr += strconv.Itoa(root.Val)
    if root.Left == nil && root.Right == nil {
        *res = append(*res, curr)
        return
    }
    
    curr += "->"
    helper(root.Left, curr, res)
    helper(root.Right, curr, res)
}