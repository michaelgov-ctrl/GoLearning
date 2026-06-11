// https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    curr := head
    for curr != nil && curr.Next != nil {
        next := curr.Next
        curr.Next = &ListNode{
            Val:  gcd(curr.Val, next.Val),
            Next: next,
        }
        curr = next
    }
    return head
}

func gcd(a, b int) int {
    if b==0 {
        return a
    }
    return gcd(b,a%b)
}

// https://leetcode.com/problems/reverse-linked-list/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var res *ListNode
    for head != nil {
        res = &ListNode{
            Val: head.Val,
            Next: res,
        }
        head = head.Next
    }

    return res
}


// https://leetcode.com/problems/add-two-numbers/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    current, carry := head, 0

    for l1 != nil || l2 != nil || carry > 0 {
        sum := carry

        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        carry = sum / 10
        current.Next = &ListNode{Val: sum % 10}
        current = current.Next
    }

    return head.Next
}



// https://leetcode.com/problems/intersection-of-two-linked-lists/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    visited := make(map[*ListNode]bool)
	for a, b := headA, headB; a != nil || b != nil; {
		if a != nil {
			if visited[a] {
				return a
			}
			visited[a] = true
			a = a.Next
		}

		if b != nil {
			if visited[b] {
				return b
			}
			visited[b] = true
			b = b.Next
		}
	}
    return nil
}


// https://leetcode.com/problems/merge-nodes-in-between-zeros/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    res, curr, sum := &ListNode{}, head.Next, 0
    tail := res
    for curr != nil {
        if curr.Val == 0 {
            tail.Next = &ListNode{Val: sum}
            tail = tail.Next
            sum = 0
        } else {
            sum += curr.Val
        }

        curr = curr.Next
    }

    return res.Next
}


// NOTE:
// [] -> []
/*
head = &ListNode{
	Val: x,
 	Next: head,
}
*/

// [] <- []
/*
tail.Next = &ListNode{Val: x}
tail = tail.Next
*/


// https://leetcode.com/problems/search-in-a-binary-search-tree/description/
// this could be improved with searching whether the current value is greater or less than 'val' since left is less and right is greater.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    return dfs(root, val)
}

func dfs(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }

    if root.Val == val {
        return root
    }

    left := dfs(root.Left, val)
    if left != nil {
        return left
    }

    right := dfs(root.Right, val)
    if right != nil {
        return right
    }

    return nil
}



// https://leetcode.com/problems/average-of-levels-in-binary-tree/description/
// this could be optimized to carry just the count & sum for each level
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    levels := [][]int{}
    flatten(root, 0, &levels)
    
    res := make([]float64, len(levels))
    for i, l := range levels {
        sum := 0
        for _, n := range l {
            sum += n
        }
        res[i] = float64(sum) / float64(len(l))
    }

    return res
}

func flatten(root *TreeNode, level int, res *[][]int) {
    if root == nil {
        return
    }

    if (len(*res) <= level) {
        *res = append(*res, []int{})
    }

    (*res)[level] = append((*res)[level], root.Val)

    flatten(root.Left, level + 1, res)
    flatten(root.Right, level + 1, res)
} 
