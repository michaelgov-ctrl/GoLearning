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
