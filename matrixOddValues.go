// https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/description/

func oddCells(m int, n int, indices [][]int) int {
    matrix := make([][]int, m)
    for i := 0; i < m; i++ {
        matrix[i] = make([]int, n)
    }
    
    for _, idx := range indices {
        for i := 0 ; i < n ; i++ {
            matrix[idx[0]][i] = matrix[idx[0]][i]+1
        }
        for i := 0 ; i < m ; i++ {
            matrix[i][idx[1]] =  matrix[i][idx[1]]+1
        }
    }

    ans := 0
    for _, row := range matrix {
        for _, v := range row {
            if v % 2 != 0 {
                ans++
            }
        }
    }

    return ans
}
