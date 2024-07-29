// https://leetcode.com/problems/find-the-encrypted-string/description/

func getEncryptedString(s string, k int) string {
    res, l := make([]byte, len(s)), len(s)
    for i, _ := range s {
        res[i] = s[wrapIndex(l, i+k)]
    }
    return string(res)
}

func wrapIndex(l, idx int) int {
    return (idx % l + l) % l
}
