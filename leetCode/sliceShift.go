// https://leetcode.com/problems/rotate-non-negative-elements/description/
func rotateElements(nums []int, k int) []int {
    positives := filterNegatives(nums)
    rotateK(positives, k)
    for i, j := 0, 0; i < len(nums) && j < len(positives); i++ {
        if nums[i] > -1 {
            nums[i] = positives[j]
            j++
        }
    }

    return nums
}

func filterNegatives(nums []int) []int {
    var positives []int
    for _, n := range nums {
        if n > -1 {
            positives = append(positives, n)
        } 
    }

    return positives
}

// https://codesignal.com/learn/courses/interview-practice-with-classic-coding-questions-in-go/lessons/advanced-slice-manipulation-techniques-in-go
func reverse(nums []int, start, end int) {
    for i, j := start, end; i < j; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
}

// shift left k steps
func rotateLeftK(nums []int, k int) {
    n := len(nums)
    if n == 0 {
        return
    }

    k = k % n

    reverse(nums, 0, k-1)
    reverse(nums, k, n-1)
    reverse(nums, 0, n-1)
}

// shift right k steps
func rotateSlice(nums []int, k int) {
    n := len(nums)
    if n == 0 {
        return
    }
    
    k = k % n // Ensure k is within the bounds of the slice length

    // Reverse entire slice
    reverse(nums, 0, n-1)
    // Reverse first k elements
    reverse(nums, 0, k-1)
    // Reverse remaining elements
    reverse(nums, k, n-1)
}
