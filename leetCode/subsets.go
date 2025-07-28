// https://leetcode.com/problems/count-number-of-maximum-bitwise-or-subsets/description/

func countMaxOrSubsets(nums []int) int {
    return NewSubsetOrTracker().Track(nums, 0, 0).MaxCount()
}

type SubsetOrTracker map[int]int

func NewSubsetOrTracker() SubsetOrTracker {
    s := make(map[int]int)
    return s
}

func (s SubsetOrTracker) Track(nums []int, i, or int) SubsetOrTracker {
    if i == len(nums) {
        s[or]++
        return s
    }
    _ = s.Track(nums, i+1, or|nums[i])
    _ = s.Track(nums, i+1, or)
    return s
}

/*
func (s SubsetOrTracker) Track(nums []int, i, or int) SubsetOrTracker {
    if i == len(nums) {
        s[or]++
        return
    }
    s.Track(nums, i+1, or|nums[i])
    s.Track(nums, i+1, or)
}
*/

func (s SubsetOrTracker) MaxCount() int {
    var max, res int
    for k, v := range s {
        if k > max {
            max, res = k, v
        }
    }
    return res
}
