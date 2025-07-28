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


// https://leetcode.com/problems/sum-of-values-at-indices-with-k-set-bits/description/
func sumIndicesWithKSetBits(nums []int, k int) int {
    var sum int
    for i, n := range nums {
        if kBitsSet(i, k) {
            sum+=n
        }
    }
    return sum
}

func kBitsSet(n, k int) bool {
    var setBits int
    for i := 0; 0 < n; i++ {
        if n&1 == 1 {
            setBits++
        }
        n >>= 1
    }
    return setBits == k
}
