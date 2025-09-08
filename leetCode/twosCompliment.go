// https://leetcode.com/problems/convert-a-number-to-hexadecimal/

// For negative numbers, num=2^32+num
// https://leetcode.com/problems/convert-a-number-to-hexadecimal/description/comments/1726347/

func toHex(num int) string {
    if num < 0 {
        num += int(math.Pow(2, 32))
    }
    return fmt.Sprintf("%x", num)
}
