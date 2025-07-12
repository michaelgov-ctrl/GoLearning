// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
// push significant bit to the left and fill the right side with head.Val
func getDecimalValue(head *ListNode) int {
    var res int
    for head != nil {
        res = (res << 1) | head.Val
        head = head.Next
    }
    return res
}
