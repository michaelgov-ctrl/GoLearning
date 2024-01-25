// https://leetcode.com/problems/left-and-right-sum-differences/description/

func leftRightDifference(nums []int) []int {
    var ans []int
    for i := 0 ; i <len(nums) ; i++ {
        absVal := abs(leftSum(nums[:i])-rightSum(nums[i:]))
        ans = append(ans, absVal)
    }

    return ans
}

func rightSum(arr []int) int {
    var sum int
    for i := 1 ; i < len(arr) ; i++ {
        sum += arr[i]
    }
    return sum
}

func leftSum(arr []int) int {
    var sum int
    for i := len(arr)-1 ; i >= 0 ; i-- {
        sum += arr[i]
    }
    return sum
}

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}
