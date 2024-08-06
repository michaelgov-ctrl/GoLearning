// https://leetcode.com/problems/design-a-stack-with-increment-operation/description/

type CustomStack struct {
    Stack []int
}


func Constructor(maxSize int) CustomStack {
    return CustomStack{Stack: make([]int, 0, maxSize)}
}


func (this *CustomStack) Push(x int)  {
    if len(this.Stack) < cap(this.Stack) {
        this.Stack = append(this.Stack, x)
    }
}


func (this *CustomStack) Pop() int {
    if len(this.Stack) == 0 {
        return -1
    }

    res := this.Stack[len(this.Stack)-1]
    this.Stack = this.Stack[:len(this.Stack)-1]
    return res
}


func (this *CustomStack) Increment(k int, val int)  {
    var l int
    if len(this.Stack) < k {
        l = len(this.Stack)
    } else {
        l = k
    }

    for i := 0; i < l; i++ {
        this.Stack[i]+=val
    }
}
