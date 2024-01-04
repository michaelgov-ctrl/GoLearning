//https://leetcode.com/problems/goal-parser-interpretation/submissions/
//a method of checking a string for expected patterns with a temporary string attempting to match pre-existing map keys
func interpret(command string) string {
    m := map[string]string{
        "G":"G",
        "()":"o",
        "(al)":"al",
    }
    var ans, tempString string
    for _, v := range command {
        tempString += string(v)
        if item, ok := m[tempString]; ok {
            ans += item
            tempString = ""
        }
    }
    return ans
}
