// https://leetcode.com/problems/decompress-run-length-encoded-list/description/
type RlePair struct {
    Frequency int
    Value int
}

type RlePairs struct {
    Pairs []RlePair
}

func (this *RlePairs) New(arr []int) {
    for i := 0 ; i < len(arr)-1 ; i+=2 {
        this.Pairs = append(this.Pairs, RlePair{Frequency: arr[i], Value: arr[i+1]})
    }
}

func (this *RlePairs) Decompress() []int {
    var res []int
    for _, v := range this.Pairs {
        temp := make([]int, v.Frequency)
        for i := range temp {
            temp[i] = v.Value
        }  
        res = append(res, temp...)
    }
    return res
}

func decompressRLElist(nums []int) []int {
    pairs := RlePairs{}
    pairs.New(nums)
    return pairs.Decompress()
}
