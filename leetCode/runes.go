// https://leetcode.com/problems/reverse-degree-of-a-string/description/

func reverseDegree(s string) int {
    var sum int
    for i, r := range s {
        sum += (i+1)*invert(r)
    }
    return sum
}

func invert(r rune) int {
    return int(rune('z')-r)+1
}
