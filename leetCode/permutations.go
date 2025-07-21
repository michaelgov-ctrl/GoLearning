
// https://leetcode.com/problems/sum-of-all-subset-xor-totals/description/

func subsetsRecursive(nums []int) [][]int {
    res := [][]int{}
    subset := []int{}

    var dfs func(int)
    dfs = func(i int) {
        // Base case: if we've processed all elements
        if i == len(nums) {
            // Create a copy of the current subset and add it to the result
            cpy := make([]int, len(subset))
            copy(cpy, subset)
            res = append(res, cpy)
            return
        }

        // Include the current element
        subset = append(subset, nums[i])
        dfs(i + 1) // Recurse with the next element

        // Exclude the current element (backtrack)
        subset = subset[:len(subset)-1]
        dfs(i + 1) // Recurse with the next element
    }

    dfs(0) // Start the recursion from the first element
    return res
}
