//https://leetcode.com/problems/matrix-diagonal-sum/

func diagonalSum(mat [][]int) int {
    if len(mat) == 1 {
        return mat[0][0]
    }

    var sum int
    for i, j, x, y := 0, 0, 0, len(mat)-1; i < len(mat[0]) && 0 <= y; i, j, x, y = i+1, j+1, x+1, y-1 {
        sum+=mat[i][j]
        sum+=mat[x][y]
    }

    l := len(mat)
    if l%2 != 0 {
        sum-=mat[(l/2)][(l/2)]
    }

    return sum
}
