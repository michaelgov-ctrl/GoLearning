func blastAlgoSearch(s1, s2 string, minSubstrLen int) string {
    for i := len(s1); i >= minSubstrLen; i-- {
        for j := 0; j <= len(s1)-i; j++ {
            subStr := s1[j:i]
            k := strings.IndexAny(s2, s1)
            if k != -1 {
                return subStr
            }
        }
    }
    return ""
}
