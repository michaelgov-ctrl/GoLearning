// https://leetcode.com/problems/sort-array-by-increasing-frequency/description/

type s struct {
    n, count int
}

type slist []s

func frequencySort(nums []int) []int {
    sl := newSlist()
    for _, n := range nums {
        sl.Increment(n)
    }
    return sl.toFrequencySortedArray()
}

func newSlist() slist {
    return []s{}
}

func (sl slist) Len() int {
    return len(sl)
}

func (sl slist) Swap(i, j int) { 
    sl[i], sl[j] = sl[j], sl[i]
}

func (sl slist) Less(i, j int) bool { 
    if sl[i].count == sl[j].count {
        return sl[i].n > sl[j].n
    }
    return sl[i].count < sl[j].count 
}

func (sl *slist) Increment(n int) {
    for i := 0; i < len(*sl); i++ {
        if (*sl)[i].n == n {
            (*sl)[i].count++
            return
        }
    }

    *sl = append(*sl, s{
        n: n,
        count: 1,
    })
}

func (sl slist) toFrequencySortedArray() []int {
    sort.Sort(sl)
    var res []int
    for _, v := range sl {
        for range v.count {
            res = append(res, v.n)
        }
    }
    return res
}

// or

func frequencySort(nums []int) []int {
    mp := make(map[int]int)
    for _, n := range nums {
        mp[n]++
    }
    sort.SliceStable(nums, func(a, b int) bool {
        if mp[nums[a]] != mp[nums[b]] {
            return mp[nums[a]] < mp[nums[b]]
        } 
        return nums[a] > nums[b]
    })
    return nums
}



// https://leetcode.com/problems/minimum-deletions-for-at-most-k-distinct-characters/description/
type charCount struct {
    char rune
    count int
}

func newCharCount(char rune, count int) charCount {
    return charCount{
        char: char,
        count: count,
    }
}

func mapToCharCountSlice(m map[rune]int) []charCount {
    res, i := make([]charCount, len(m)), 0
    for k, v := range m {
        res[i] = newCharCount(k, v)
        i++
    }

    return res
}

func stringToCharCountSlice(s string) []charCount {
    m := make(map[rune]int)
    for _, r := range s {
        m[r]++
    }

    return mapToCharCountSlice(m)
}

func minDeletion(s string, k int) int {
    charCounts := stringToCharCountSlice(s)
    slices.SortFunc(charCounts, func(a, b charCount) int {
        return cmp.Compare(a.count, b.count)
    })

    res := 0
    for i := 0; i < len(charCounts)-k; i++ {
        res += charCounts[i].count
    }
    
    return res
}




// https://leetcode.com/problems/find-x-sum-of-all-k-long-subarrays-i/description/
func findXSum(nums []int, k int, x int) []int {
    res := make([]int, len(nums)-k+1)
    for i := 0; i <= len(nums)-k; i++ {
        res[i] = xSum(nums[i:i+k], x)
    }
    return res
}

func sum(nums []int) int {
    var res int
    for _, n := range nums {
        res+=n
    }
    return res
}

type Freq struct {
    Val, Count int
}

func newFreq(val, count int) Freq {
    return Freq{
        Val: val,
        Count: count,
    }
}

func xSum(nums []int, x int) int {
    freqMap := make(map[int]int)
    for _, n := range nums {
        freqMap[n]++
    }

    freqArr, i := make([]Freq, len(freqMap)), 0
    for k, v := range freqMap {
        freqArr[i] = newFreq(k, v)
        i++
    }

    if len(freqArr) <= x {
        var res int
        for _, f := range freqArr {
            res += f.Val * f.Count
        }
        return res
    }

    slices.SortFunc(freqArr, func(a, b Freq) int {
        if n := cmp.Compare(b.Count, a.Count); n != 0 {
            return n
        }
        return cmp.Compare(b.Val, a.Val)
    })

    var res int
    for i := 0; i < x; i++ {
        res += freqArr[i].Val * freqArr[i].Count
    }

    return res
}
