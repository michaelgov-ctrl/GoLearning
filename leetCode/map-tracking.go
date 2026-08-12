// https://leetcode.com/problems/design-underground-system/

type Customer struct {
    start int
    station string
}

type Trips struct {
    total int
    count int
}

type UndergroundSystem struct {
    tracker map[int]Customer
    stationVisits map[string]map[string]Trips
}

func Constructor() UndergroundSystem {
    return UndergroundSystem{
        tracker: make(map[int]Customer),
        stationVisits: make(map[string]map[string]Trips),
    }
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
    this.tracker[id] = Customer{
        start: t,
        station: stationName,
    }
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
    customer := this.tracker[id]
    delete(this.tracker, id)

    if _, ok := this.stationVisits[customer.station]; !ok {
        this.stationVisits[customer.station] = make(map[string]Trips)
    }

    trips := this.stationVisits[customer.station][stationName]
    trips.total += t - customer.start
    trips.count++
    this.stationVisits[customer.station][stationName] = trips
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    trips := this.stationVisits[startStation][endStation]
    return float64(trips.total) / float64(trips.count)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */


// https://leetcode.com/problems/rings-and-rods/description/
func countPoints(rings string) int {
    runes, rods := []rune(rings), make(map[rune]map[rune]struct{})
    for i := 0; i < len(rings); i += 2 {
        if rods[runes[i+1]] == nil {
            rods[runes[i+1]] = make(map[rune]struct{})
        }

        rods[runes[i+1]][runes[i]] = struct{}{}
    }
    
    var res int
    for _, v := range rods {
        if len(v) == 3 {
            res++
        }
    }

    return res
}
