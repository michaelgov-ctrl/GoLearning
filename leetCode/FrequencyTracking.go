// https://leetcode.com/problems/length-of-longest-subarray-with-at-most-k-frequency/description
func maxSubarrayLength(nums []int, k int) int {
    res, freq, i := 0, make(map[int]int), 0
    for j := 0; j < len(nums); j++ {
        freq[nums[j]]++

        for freq[nums[j]] > k {
            freq[nums[i]]--
            i++
        }

        res = max(res, j-i+1)
    }

    return res
}
