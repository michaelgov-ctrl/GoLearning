func splitString(s string) (letters string, digit int) {
    var str, d []rune
    for _, r := range s {
        switch {
        case r >= 'A' && r <= 'Z':
            str = append(str, r)
        case r >= 'a' && r <= 'z':
            str = append(str, r)
        case r >= '0' && r <= '9':
            d = append(d, r)
        }
    }
    digit, err := strconv.Atoi(string(d))
    if err != nil {
        return 
    }
    return string(str), digit
}
