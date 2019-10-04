package stats

import "fmt"

type Stat struct {
	SectionHits map[string]int
	Failures    map[string]int
}

type StatList struct {
	FirstStatIndex int
	BenchmarkTime  int64
	Stats          [2]Stat
}

func NewStat(sectionHits map[string]int, failures map[string]int) Stat {
	return Stat{
		SectionHits: sectionHits,
		Failures:    failures,
	}
}

func NewStatList() StatList {
	stats := [2]Stat{
		NewStat(make(map[string]int), make(map[string]int)),
		NewStat(make(map[string]int), make(map[string]int)),
	}

	return StatList{
		FirstStatIndex: 0,
		BenchmarkTime:  0,
		Stats:          stats,
	}
}

func (sl *StatList) UpdateStatList(section string, status int, date int64) string {
	statString := ""
	if sl.BenchmarkTime == 0 {
		sl.BenchmarkTime = date
	}

	firstBucketEnd := sl.BenchmarkTime + 10
	secondBucketEnd := sl.BenchmarkTime + 20

	if date >= sl.BenchmarkTime && date < firstBucketEnd {
		if sl.FirstStatIndex == 0 {
			sl.Stats[0].updateStat(section, status)
		} else {
			sl.Stats[1].updateStat(section, status)
		}
	} else if date >= firstBucketEnd && date < secondBucketEnd {
		if sl.FirstStatIndex == 0 {
			sl.Stats[1].updateStat(section, status)
		} else {
			sl.Stats[0].updateStat(section, status)
		}
	} else if date >= secondBucketEnd {
		if sl.FirstStatIndex == 0 {
			statString = sl.Stats[0].Flush()
			sl.FirstStatIndex = 1
			sl.BenchmarkTime = sl.BenchmarkTime + 10
			sl.Stats[0].updateStat(section, status)
		} else {
			statString = sl.Stats[1].Flush()
			sl.FirstStatIndex = 0
			sl.BenchmarkTime = sl.BenchmarkTime + 10
			sl.Stats[1].updateStat(section, status)
		}
	}

	return statString
}

func (sl *StatList) FlushAll() string {
	if sl.FirstStatIndex == 0 {
		return sl.Stats[0].Flush() + sl.Stats[1].Flush()
	} else {
		return sl.Stats[1].Flush() + sl.Stats[0].Flush()
	}
}

func (s *Stat) Flush() string {
	section := s.maxHitSection()
	hits := s.SectionHits[section]
	s.reset()
	return fmt.Sprintf("Section with most hits: %s with %d hits.\n", section, hits)
}

func (s *Stat) maxHitSection() string {
	maxHitSection := ""
	for k, v := range s.SectionHits {
		if v >= s.SectionHits[maxHitSection] {
			maxHitSection = k
		}
	}
	return maxHitSection
}

func (s *Stat) reset() {
	s.SectionHits = make(map[string]int)
	s.Failures = make(map[string]int)
}

func (s *Stat) updateStat(section string, status int) {
	s.SectionHits[section]++

	if status != 200 {
		s.Failures[section]++
	}
}
