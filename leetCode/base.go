
// https://leetcode.com/problems/hexadecimal-and-hexatrigesimal-conversion/description/
// hexadeximal(base 16) & hexatrigesimal(base 36)
func concatHex36(n int) string {
    return strings.ToUpper(strconv.FormatInt(int64(n*n), 16) + strconv.FormatInt(int64(n*n*n), 36))
}
