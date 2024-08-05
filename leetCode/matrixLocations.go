// https://leetcode.com/problems/design-neighbor-sum-service/description/

type location struct {
    Row int
    Col int
}

type neighborSum struct {
    ValueLocation map[int]location
    LocationValue map[location]int
}


func Constructor(grid [][]int) neighborSum {
    n := neighborSum{
        ValueLocation: make(map[int]location),
        LocationValue: make(map[location]int),
    }

    for y, row := range grid {
        for x, v := range row {
            l := location{Row: y, Col: x}
            n.ValueLocation[v]=l
            n.LocationValue[l]=v
        }
    }

    fmt.Println(n)
    return n
}


func (this *neighborSum) AdjacentSum(value int) int {
    root := this.ValueLocation[value]
    up := location{Row: root.Row+1, Col: root.Col}
    down := location{Row: root.Row-1, Col: root.Col}
    right := location{Row: root.Row, Col: root.Col+1}
    left :=  location{Row: root.Row, Col: root.Col-1}

    return this.LocationValue[up] + this.LocationValue[down] + this.LocationValue[right] + this.LocationValue[left]
}


func (this *neighborSum) DiagonalSum(value int) int {
    root := this.ValueLocation[value]
    upRight := location{Row: root.Row+1, Col: root.Col+1}
    upLeft := location{Row: root.Row+1, Col: root.Col-1}
    downRight := location{Row: root.Row-1, Col: root.Col+1}
    downLeft :=  location{Row: root.Row-1, Col: root.Col-1}

    return this.LocationValue[upRight] + this.LocationValue[upLeft] + this.LocationValue[downRight] + this.LocationValue[downLeft]
}


/**
 * Your neighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */
