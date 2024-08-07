// https://leetcode.com/problems/check-if-number-is-a-sum-of-powers-of-three/description/

func checkPowersOfThree(n int) bool {
    for n > 0 {
        if 1 < n%3 {
            return false
        }
        n /= 3
    }

    return true
}

func asBase3(n int) int {
    res, factor := 0, 1
    for n > 0 {
        res += n % 3 * factor
        n /= 3
        factor *= 10
    }
    return res
}
