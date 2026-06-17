// https://leetcode.com/problems/find-common-characters/description/

// take characters and frequencies from shortest word,
// then min on character occurences for each word.
func commonChars(words []string) []string {
    res, shortest := []string{}, shortestWordFreq(words)
Outer:
    for k, v := range shortest {
        minimum, char := v, string(k)
        for _, w := range words {
            minimum = min(minimum, strings.Count(w, char))
            if minimum == 0 {
                continue Outer
            }
        }

        for range minimum {
            res = append(res, char)
        }    
    }

    return res
}

func shortestWordFreq(words []string) map[rune]int {
    shortest, min := "", math.MaxInt
    for _, w := range words {
        if len(w) < min {
            shortest, min = w, len(w)
        }
    }

    res := make(map[rune]int)
    for _, r := range shortest {
        res[r]++
    }

    return res
}
