// https://leetcode.com/problems/excel-sheet-column-title/description/

func convertToTitle(columnNumber int) string {
    n, res := columnNumber, []byte{}
    for n > 0 {
        n--
        r := n % 26
        res = append([]byte{byte('A'+r)}, res...)
        n/=26 
    }

    return string(res)
}

// inverse
// https://leetcode.com/problems/excel-sheet-column-number/description/

func titleToNumber(columnTitle string) int {
    var res int
    for _, r := range columnTitle {
        res *= 26
        res += int(r - 'A' + 1)
    }
    return res
}
