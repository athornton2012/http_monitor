package stats

import "fmt"

type RollingTrafficList struct {
	TrafficList      []int //assuming that there is no more than 10 seconds of clock skew
	StartIndex       int
	LatestDate       int64
	HighTrafficAlert bool
	Limit            int
	TotalTraffic     int
	WindowSize       int
}

func NewRollingTrafficList(limit, windowSize int) RollingTrafficList {
	return RollingTrafficList{
		TrafficList:      make([]int, windowSize, windowSize),
		StartIndex:       0,
		LatestDate:       0,
		HighTrafficAlert: false,
		Limit:            limit,
		TotalTraffic:     0,
		WindowSize:       windowSize,
	}
}

func (r *RollingTrafficList) HandleLog(date int64) string {
	if r.LatestDate == 0 {
		r.LatestDate = date
	}

	offset := int(date - r.LatestDate)

	if date <= r.LatestDate {
		index := r.StartIndex + offset
		if index < 0 { //loop around to the end of the array
			index = r.WindowSize + index
		}

		r.TrafficList[index]++
		r.TotalTraffic++
	} else {
		sum := 0
		for i := 1; i <= offset; i++ {
			indexToOverwrite := (r.StartIndex + i) % r.WindowSize
			sum += r.TrafficList[indexToOverwrite]
			r.TrafficList[indexToOverwrite] = 0
		}
		r.StartIndex = (r.StartIndex + offset) % r.WindowSize
		r.TrafficList[r.StartIndex]++
		r.TotalTraffic = r.TotalTraffic - sum + 1
		r.LatestDate = date
	}

	averageTraffic := r.TotalTraffic / r.WindowSize
	if averageTraffic >= r.Limit {
		if !r.HighTrafficAlert {
			r.HighTrafficAlert = true
			return "HIGH TRAFFIC ALERT\n"
		}
	} else { //averageTraffic is under the limit
		if r.HighTrafficAlert {
			r.HighTrafficAlert = false
			return fmt.Sprintf("NORMAL TRAFFIC RECOVERED BETWEEN %d and %d\n", r.LatestDate-1, r.LatestDate+1)
		}
	}

	return ""
}
