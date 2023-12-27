func intToSlice(num int) []int {
    arr := []int{}
    for num > 0 {
        arr = append(arr, num%10)
        num = int(num / 10)
    }
    return arr
}
