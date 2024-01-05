// https://leetcode.com/problems/maximum-number-of-coins-you-can-get/description/
func maxCoins(piles []int) int {
    slices.Sort(piles)
    var ans int
    for i := len(piles)-2 ; len(piles)/3 <= i ; i-=2 {
        ans += piles[i]
    }
    return ans
}
