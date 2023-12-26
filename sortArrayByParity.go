func sortArrayByParity(nums []int) []int {
    ans := make([]int, len(nums))
    start, end := 0, len(nums)-1
    for _, v := range nums {
        if v & 0x1 == 0 {
            ans[start] = v
            start++
        } else {
            ans[end] = v
            end--
        }
    }
    return ans
}
