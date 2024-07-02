// https://leetcode.com/problems/clear-digits/
// this contains an example of deleting an item from an array/slice: deleteElement()

func clearDigits(s string) string {
    foundDigit := true
    for foundDigit {
        foundDigit, s = removeDigitAndLeftAlpha(s)
    }

    return s
}

func removeDigitAndLeftAlpha(s string) (bool, string) {
    runes := []rune(s)
    for i, v := range runes {
        if isDigit(v) {
            runes = deleteElement(runes, i)
            return searchLeftAndRemoveAlpha(runes, i-1)
        }
    }

    return false, string(runes)
}

func searchLeftAndRemoveAlpha(runes []rune, i int) (bool, string) {
    for ; 0 <= i; i-- {
        if !isDigit(runes[i]) {
            return true, string(deleteElement(runes, i))
        }
    }

    return false, string(runes)
}

func deleteElement(runes []rune, i int) []rune {
    if i < len(runes)-1 {
        return append(runes[:i], runes[i+1:]...)
    }

    return runes[:i]
}

func isDigit(r rune) bool {
    return rune('0') <= r && r <= rune('9')
}
