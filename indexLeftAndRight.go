// index from left and right
// https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/description/

func minPairSum(nums []int) int {
    slices.Sort(nums)

    var max int
    for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
        if nums[i]+nums[j] > max {
            max = nums[i]+nums[j]
        }
    }

    return max
}

func minPairSum(nums []int) int {
    sort.Ints(nums)
    max := 0

    for i := 0; i < len(nums); i++ {
        if nums[i] + nums[len(nums)-1-i] > max {
            max =  nums[i] + nums[len(nums)-1-i]
        }
    }

    return max
    
}
