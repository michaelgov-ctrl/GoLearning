//https://leetcode.com/problems/create-target-array-in-the-given-order/solutions/4501228/go-0ms-runtime/

func shiftArrayValuesRight(idx int, arr []int ) []int {
    for i := len(arr)-1 ; i > idx ; i-- {
        arr[i] = arr[i-1]
    }
    return arr
}

func createTargetArray(nums []int, index []int) []int {
    ans := make([]int, len(nums))
    
    for i := range ans {
        ans[i] = -1
    }
    
    for i := 0 ; i < len(ans) ; i++ {
        if ans[index[i]] == -1 {
            ans[index[i]] = nums[i]
        } else {
            ans = shiftArrayValuesRight(index[i], ans)
            ans[index[i]] = nums[i]
        }
    }

    return ans
}
