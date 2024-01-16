
// https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/description/
func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
    var fw, sw, tw int

    //convert first word to int with a=0, b=1, c=1,...
    base := 1
    for i := len(firstWord)-1 ; i >= 0 ; i-- {
        fw += (int(firstWord[i])-97)*base
        base = base*10
    }

    //convert second word to int with a=0, b=1, c=1,...
    base = 1
    for i := len(secondWord)-1 ; i >= 0 ; i-- {
        sw += (int(secondWord[i])-97)*base
        base = base*10
    }

    //convert target word to int with a=0, b=1, c=1,...
    base = 1
    for i := len(targetWord)-1 ; i >= 0 ; i-- {
        tw += (int(targetWord[i])-97)*base
        base = base*10
    }
    
    if fw+sw == tw {
        return true
    }

    return false
}
