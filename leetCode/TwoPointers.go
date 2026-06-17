// https://leetcode.com/problems/is-subsequence/description/
func isSubsequence(s string, t string) bool {
    var i, j int
    for i < len(s) && j < len(t) {
        if s[i] == t[j] {
            i++
        }
        j++
    }
    
    return i == len(s)
}

// https://leetcode.com/problems/long-pressed-name/description/
func isLongPressedName(name string, typed string) bool {
    var i, j int
    for j < len(typed) {
        if i < len(name) && typed[j] == name[i] {
            i++
            j++
        } else if j > 0 && typed[j-1] == typed[j] {
            j++
        } else {
            return false
        }
    }
    
    return i == len(name)
}
