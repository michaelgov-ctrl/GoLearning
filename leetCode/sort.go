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
