// https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-i/description/

type CharOccurence struct {
    Char rune
    Occurence int
}

type ByOccurence []CharOccurence
func (a ByOccurence) Len() int           { return len(a) }
func (a ByOccurence) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByOccurence) Less(i, j int) bool { return a[i].Occurence < a[j].Occurence }

func maxDifference(s string) int {
    charCount := make(map[rune]int)
    for _, r := range s {
        charCount[r]++
    }

    var evens, odds []CharOccurence
    for k, v := range charCount {
        if v % 2 == 0 {
            evens = append(evens, CharOccurence{
                Char: k,
                Occurence: v,
            })
        } else {
            odds = append(odds, CharOccurence{
                Char: k,
                Occurence: v,
            })
        }
    }

    sort.Sort(ByOccurence(evens))
    sort.Sort(ByOccurence(odds))

    leastOddMostEven := odds[0].Occurence - evens[len(evens)-1].Occurence
    mostOddLeastEven := odds[len(odds)-1].Occurence - evens[0].Occurence
    if leastOddMostEven > mostOddLeastEven {
        return leastOddMostEven
    }
    
    return mostOddLeastEven
}
