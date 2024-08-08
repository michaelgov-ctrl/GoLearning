// https://leetcode.com/problems/evaluate-the-bracket-pairs-of-a-string/description/

func evaluate(s string, knowledge [][]string) string {
    var sb,key strings.Builder    
    keys := make(map[string]string)

    for _,v := range knowledge {
        keys[v[0]] = v[1]
    }
    
    for i := 0; i < len(s); i++ {
        if s[i] == '('  {  
            for i++ ; s[i] != ')'; i++ {
                key.WriteByte(s[i])
            }

            if _, ok := keys[key.String()]; !ok {
                sb.WriteByte('?')
            } else {
                sb.WriteString(keys[key.String()])
            }

            key.Reset()
            continue
        }

        sb.WriteByte(s[i])
    }
         
    return sb.String()
}
