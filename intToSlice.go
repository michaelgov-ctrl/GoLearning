func intToSlice(num int) []int {
    base, n := 0, num
    for n > 0 {
        n = n/10
        base++
    }
    base, arr := 0, make([]int, base)
    for num > 0 {
        arr[len(arr)-(base+1)] = num%10
        num = int(num / 10)
        base++
    }
    return arr
}

//the opposite (going from slice to single integer)
//https://stackoverflow.com/questions/44729587/join-digits-from-integer-array-into-one-number-in-golang

func sliceToInt(s []int) int {
    res := 0
    op := 1
    for i := len(s) - 1; i >= 0; i-- {
        res += s[i] * op
        op *= 10
    }
    return res
}
