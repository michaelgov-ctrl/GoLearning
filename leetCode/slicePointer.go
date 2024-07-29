// https://leetcode.com/problems/distribute-elements-into-two-arrays-i/description/

func resultArray(nums []int) []int {
    arr1, arr2 := []int{nums[0]}, []int{nums[1]}
    for i := 2; i < len(nums); i++ {
        appendVal(&arr1, &arr2, arr1[len(arr1)-1], arr2[len(arr2)-1], nums[i])
    }

    return append(arr1, arr2...)
}

func appendVal(arr1, arr2 *[]int, elem1, elem2, val int) {
    if elem1 > elem2 {
        *arr1 = append(*arr1, val)
    } else {
        *arr2 = append(*arr2, val)
    }
}
