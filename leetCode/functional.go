
// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
    runes, duplicateCheck := []rune(s), func(r []rune) bool {
        seen := make(map[rune]bool)
        for _, v := range r {
            if seen[v] {
                return false
            }
            seen[v] = true
        }
        return true
    }

    for k := len(runes)-1; k >= 0; k-- {
        if slidingWindowFunc(runes, k, duplicateCheck) {
            return k+1
        }
    }

    return -1
}

func slidingWindowFunc(r []rune, k int, f func([]rune) bool) bool {
    for i := 0; i < len(r)-k; i++ {
        if (f(r[i:k+i+1])) {
            return true
        }
    }
    return false
}

// much simpler solution than above
func lengthOfLongestSubstring(s string) int {
    res, tracking := 0, make([]int, 128) // # of diff chars
    for start, end := 0, 0; end < len(s); end++ {
        curr := s[end]
        if tracking[curr] > start {
            start = tracking[curr]
        }

        if end-start+1 > res {
            res = end-start+1
        }

        tracking[curr] = end+1
    }

    return res
}
