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
