// https://leetcode.com/problems/kth-distinct-string-in-an-array/

func kthDistinct(arr []string, k int) string {
    m := make(map[string]int)
    for _, str := range arr {
        m[str]++
    }

    for _, str := range arr {
        if m[str] == 1 {
            k--
            if k == 0 {
                return str
            }
        }
    }

    return ""
}
