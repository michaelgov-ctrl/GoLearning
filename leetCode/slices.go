// https://leetcode.com/problems/arithmetic-subarrays/description/
func checkArithmeticSubarrays(nums, lefts, rights []int) []bool {
    if len(lefts) != len(rights) {
        panic("this aint gonna work")
    }

    res := make([]bool, len(lefts))
Outer:
    for i := 0; i < len(lefts); i++ {
        l, r := lefts[i], rights[i]
        if r-l+1 < 2 {
            panic("this aint gonna compare")
        }

        cp := append([]int{}, nums[l:r+1]...)
        slices.Sort(cp)
        
        diff := cp[1]-cp[0]
        for i := 0; i < len(cp)-1; i++ {
            if cp[i+1]-cp[i] != diff {
                continue Outer
            }
        }

        res[i] = true
    }

    return res
}

// or

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
    ans := make([]bool, len(l))
    for i := 0; i < len(l); i++ {
        ans[i] = checkArr(append([]int{}, nums[l[i]:r[i]+1]...))
    }
    return ans
}

func checkArr(nums []int) bool {
    sort.Ints(nums)
    for i := 2; i < len(nums); i++ {
        if nums[i-1] - nums[i] != nums[0] - nums[1] {
            return false
        }
    }
    return true
}
