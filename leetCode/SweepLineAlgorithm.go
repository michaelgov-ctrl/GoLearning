// https://leetcode.com/problems/maximum-population-year/description/

// sweep line algorithm

type terminus int

const (
    start terminus = iota
    end
)

type event struct {
    year int
    terminus terminus
}

func newEvent(year int, terminus terminus) event {
    return event{
        year: year,
        terminus: terminus,
    }
}

type events []event

func eventsFromLogs(logs [][]int) events {
    var events events
    for _, lifespan := range logs {
        events = append(events, newEvent(lifespan[0], start))
        events = append(events, newEvent(lifespan[1], end))
    }

    return events
}

func (e events) sort() events {
    sort.Slice(e, func(i, j int) bool {
        if e[i].year == e[j].year {
            return e[i].terminus > e[j].terminus
        }

        return e[i].year < e[j].year
    })

    return e
}

func (e events) maxOverlap() (overlap int, year int) {
    var curr int
    for _, event := range e {
        change := 1 // with each event we're either adding
        if event.terminus == end {
            change = -1 // or removing an overlap
        }

        curr += change

        if curr > overlap {
            overlap, year = curr, event.year
        }
    }

    return
}

func maximumPopulation(logs [][]int) int {
    _, year := eventsFromLogs(logs).sort().maxOverlap()
    return year
}
