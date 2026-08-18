
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


// https://leetcode.com/problems/find-the-largest-almost-missing-integer/description
func largestInteger(nums []int, k int) int {
    totalFreq := make(map[int]int)
    for i := 0; i <= len(nums)-k; i++ {
        tempFreq := make(map[int]struct{})
        for _, n := range nums[i:k+i] {
            tempFreq[n] = struct{}{}
        }

        for k := range tempFreq {
            totalFreq[k]++
        }
    }

    var candidates []int
    for k, v := range totalFreq {
        if v == 1 {
            candidates = append(candidates, k)
        }
    }

    if len(candidates) == 0 {
        return -1
    }

    return slices.Max(candidates)
}

// or

func largestInteger(nums []int, k int) int {
    if k > len(nums) {
		return -1
	}

    m := make(map[int]int)
	for i := 0; i <= len(nums)-k; i++ {
		window, temp := nums[i:i+k], make(map[int]bool)        
        for _, n := range window {
            temp[n] = true
        }

        for k := range temp {
            m[k]++
        }
    }

    res := -1
    for k, v := range m {
        if v == 1 && k > res {
            res = k
        }
    }

    return res
}
