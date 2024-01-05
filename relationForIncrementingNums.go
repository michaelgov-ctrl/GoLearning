// https://leetcode.com/problems/maximum-sum-with-exactly-k-elements/description/
func maximizeSum(nums []int, k int) int {
    var target int
    for _, v := range nums {
        if target < v {
            target = v
        }
    }
    return target*k+((k*(k-1))/2)
  /*
    relationship for incrementing after first digit broken out
    result := k * largest
	  result += (k * (k - 1)) / 2
  */
}
