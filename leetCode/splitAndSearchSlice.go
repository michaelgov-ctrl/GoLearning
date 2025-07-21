// https://leetcode.com/problems/maximum-score-after-splitting-a-string/

func maxScore(s string) int {
    arr, maxScore := []rune(s), 0
    for i := 1; i < len(arr); i++ {
        tempScore := searchLeftZeros(arr[:i]) + searchRightOnes(arr[i:])
        if maxScore < tempScore {
            maxScore = tempScore
        }
    }

    return maxScore
}

func searchLeftZeros(arr []rune) int {
    var res int
    for i := len(arr)-1; 0 <= i; i-- {
        if arr[i] == '0' {
            res++
        }
    }

    return res
}

func searchRightOnes(arr []rune) int {
    var res int
    for i := 0; i < len(arr); i++ {
        if arr[i] == '1' {
            res++
        }
    }

    return res
}

// https://leetcode.com/problems/queries-on-a-permutation-with-key/
// multi append around deletion
func processQueries(queries []int, m int) []int {
    p, res := generateP(m), make([]int, len(queries))
    for i := 0; i < len(res); i++ {
        idx := slices.Index(p, queries[i])
        res[i] = idx
        if idx < len(p) {
            temp := append(p[:idx], p[idx+1:]...)
            p = append([]int{queries[i]}, temp...)
        }
    }
    return res
}

func generateP(m int) (p []int) {
    p = make([]int, m)
    for i := 0; i < len(p); i++ {
        p[i] = i+1
    }
    return
}
