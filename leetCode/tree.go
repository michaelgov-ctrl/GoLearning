// https://leetcode.com/problems/evaluate-boolean-binary-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
    switch root.Val {
    case 0:
        return false
    case 2:
        return evaluateTree(root.Left) || evaluateTree(root.Right)
    case 3:
        return evaluateTree(root.Left) && evaluateTree(root.Right)
    default:
        return true
    }
}
