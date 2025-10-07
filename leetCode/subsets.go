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

// https://leetcode.com/problems/subsets/description/
func subsets(nums []int) [][]int {
	res, curr := [][]int{}, []int{}

    var backtrack func(start int)
	backtrack = func(start int) {
		temp := make([]int, len(curr))
		copy(temp, curr)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			curr = append(curr, nums[i])
			backtrack(i + 1)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0)
	return res
}
