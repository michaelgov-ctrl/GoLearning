
// https://leetcode.com/problems/license-key-formatting/description
// https://go.dev/wiki/SliceTricks#insert
func licenseKeyFormatting(s string, k int) string {
    cleaned := strings.Map(func(r rune) rune {
        if r == '-' {
            return -1
        }
        
        if unicode.IsLower(r) {
            return unicode.ToUpper(r)
        }
        
        return r
    }, s)

    runes, counter := []rune(cleaned), 0
    for i := len(runes)-1; i > 0; i-- {
        counter++
        if counter == k {
            runes = append(runes, '\x00')
            copy(runes[i+1:], runes[i:])
            runes[i] = '-'
            counter = 0
        }
    }

    return string(runes)
}
