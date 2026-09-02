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



// https://leetcode.com/problems/binary-tree-preorder-traversal/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    var res []int
    
    var preorder func(node *TreeNode)
    preorder = func(node *TreeNode) {
        if node == nil {
            return
        }

        res = append(res, node.Val)
        preorder(node.Left)
        preorder(node.Right)
    }

    preorder(root)

    return res
}




// https://leetcode.com/problems/binary-tree-postorder-traversal/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    var res []int

    var postorder func(node *TreeNode)
    postorder = func(node *TreeNode) {
        if node == nil {
            return
        }

        postorder(node.Left)
        postorder(node.Right)

        res = append(res, node.Val)
    }

    postorder(root)

    return res
}

