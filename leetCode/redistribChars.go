// https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/description/

func makeEqual(words []string) bool {
    m := make(map[rune]int)
    for _, word := range words {
        for _, r := range word {
            m[r]++
        }
    }

    l := len(words)
    for _, v := range m {
        if v % l != 0 {
            return false
        }
    }

    return true
}
