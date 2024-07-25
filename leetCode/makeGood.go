// https://leetcode.com/problems/make-the-string-great/description/

func makeGood(s string) string {
    runes, found := []rune(s), true
    for found {
        runes, found = removedAdjSameChars(runes)
    }
    return string(runes)
}

func removedAdjSameChars(runes []rune) ([]rune, bool) {
    for i := 0; i < len(runes)-1; i++ {
        if isSame(runes[i], runes[i+1]) {
            return remove2(runes, i), true
        }
    }
    return runes, false
}

func isSame(r1, r2 rune) bool {
    if (r1-32) == r2 || (r1+32) == r2 {
        return true
    }
    return false
}

func remove2(slice []rune, s int) []rune {
    return append(slice[:s], slice[s+2:]...)
}
