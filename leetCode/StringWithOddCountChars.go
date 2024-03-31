// https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/description/

func generateTheString(n int) string {
    if n % 2 == 0 {
        return strings.Repeat("a", (n-1)) + "b"
    }
    return strings.Repeat("a", n)
}
