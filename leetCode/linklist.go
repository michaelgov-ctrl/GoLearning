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
