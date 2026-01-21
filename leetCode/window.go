
// https://leetcode.com/problems/existence-of-a-substring-in-a-string-and-its-reverse/description/
// window of length 2 across string (chars[i-1:i+1)
func isSubstringPresent(s string) bool {
    rev, chars := reverse(s), []rune(s)
    for i := 1; i < len(chars); i++ {
        subStr := string(chars[i-1:i+1])
        if strings.Contains(rev, subStr) {
            return true
        }
    }

    return false
}

func reverse(s string) string {
    res := []rune(s)
    for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }

    return string(res)
}
