// https://leetcode.com/problems/count-commas-in-range-ii/description/
// https://leetcode.com/problems/count-commas-in-range-ii/solutions/7649010/math-trick-count-commas-by-digit-thresho-nxqt/

func countCommas(n int64) int64 {
    res, t := int64(0), int64(1000)
    for t <= n {
        res += n-t+1
        t *= 1000
    }
    return res
}
