
// https://leetcode.com/problems/remove-all-occurrences-of-a-substring/description/
func removeOccurrences(s string, part string) string {
    idx := strings.Index(s, part)
    for idx != -1 {
        before, after, found := strings.Cut(s, part)
        if !found {
            break
        }
        
        s = before+after
        idx = strings.Index(s, part)
    }

    return s
}

// or
func removeOccurrences(s string, part string) string {
    for strings.Index(s, part) != -1 {
        before, after, found := strings.Cut(s, part)
        if !found {
            break
        }
        
        s = before+after
    }

    return s
}

// or
func removeOccurrences(s string, part string) string {
    for {
        before, after, found := strings.Cut(s, part)
        if !found {
            break
        }
        
        s = before+after
    }

    return s
}
