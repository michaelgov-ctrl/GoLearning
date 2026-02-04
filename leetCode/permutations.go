
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



// https://leetcode.com/problems/maximum-sum-of-three-numbers-divisible-by-three/description/
func maximumSum(nums []int) int {
    var max int
    var current []int
    used := make([]bool, len(nums))

    var dfs func(current []int)
    dfs = func(current []int) {
        if len(current) == 3 {
            temp := sum(current)
            if temp % 3 == 0 && temp > max {
                max = temp
            }

            return
        }

        for i := 0; i < len(nums); i++ {
            if !used[i] {
                used[i] = true
                current = append(current, nums[i])
                dfs(current)

                current = current[:len(current)-1]
                used[i] = false
            }
        }
    }

    dfs(current)

    return max
}

func sum(nums []int) int {
    var sum int
    for _, n := range nums {
        sum += n
    }

    return sum
}

func generateKPermutations(nums []int, k int) [][]int {
    var results [][]int
    var current []int
    used := make([]bool, len(nums))

    var dfs func(current []int)
    dfs = func(current []int) {
        if len(current) == k {
            temp := make([]int, k)
            copy(temp, current)
            results = append(results, temp)
            return
        }

        for i := 0; i < len(nums); i++ {
            if !used[i] {
                used[i] = true
                current = append(current, nums[i])
                dfs(current)

                current = current[:len(current)-1]
                used[i] = false
            }
        }
    }

    dfs(current)

    return results
}
