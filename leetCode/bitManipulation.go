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

// https://leetcode.com/problems/number-complement/description/
// Invert bits and trim significant digits for accurate representation with mask
// The binary representation of 5 is 101 (no leading zero bits), and its complement is 010 so 2.
func findComplement(num int) int {
    numBits := bits.Len(uint(num))
    if numBits == 0 {
		return 0
	}

    mask := (1 << numBits) - 1
    return num ^ mask
}

// https://leetcode.com/problems/number-of-1-bits/description/
// count set bits of an int
func hammingWeight(n int) int {
    var res int
    for i := 0; i < strconv.IntSize; i++ {
        if n & 1 == 1 {
            res++
        }
        n >>= 1
    }
    return res
}

// more simple version of above is to move "downwards"
func hammingWeight(n int) int {
    var res int
    for n != 0 {
        res += n & 1
        n >>= 1
    }
    return res
}
