package stats_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/athornton2012/http_monitor/stats"
)

var _ = Describe("StatList", func() {
	Describe("#UpdateStatList", func() {
		var (
			s                StatList
			stat1SectionHits map[string]int
			stat2SectionHits map[string]int
			stat1Failures    map[string]int
			stat2Failures    map[string]int
		)

		BeforeEach(func() {
			stat1SectionHits = map[string]int{"some section": 1, "some other section": 6}
			stat1Failures = map[string]int{"some section": 1, "some other section": 1}

			stat2SectionHits = map[string]int{"some section": 7, "some other section": 5}
			stat2Failures = map[string]int{"some section": 1, "some other section": 2}
		})

		Context("The first stat index is 0", func() {
			Context("The time is more than 20 seconds from the benchmark time", func() {
				It("returns a 10 second window update, clears the first index, increments the second and swaps the index of the first bucket", func() {
					s = StatList{
						FirstStatIndex: 0,
						BenchmarkTime:  1000,
						Stats:          [2]Stat{NewStat(stat1SectionHits, stat1Failures), NewStat(stat2SectionHits, stat2Failures)},
					}

					Expect(s.UpdateStatList("some section", 200, 1023)).To(Equal("Section with most hits: some other section with 6 hits.\n"))
					Expect(s.BenchmarkTime).To(Equal(int64(1010)))
					Expect(s.FirstStatIndex).To(Equal(1))
					Expect(s.Stats[0].SectionHits).To(Equal(map[string]int{"some section": 1}))
					Expect(s.Stats[0].Failures).To(Equal(map[string]int{}))
					Expect(s.Stats[1].SectionHits).To(Equal(map[string]int{"some section": 7, "some other section": 5}))
					Expect(s.Stats[1].Failures).To(Equal(map[string]int{"some section": 1, "some other section": 2}))
				})
			})

			Context("The time is 10-20 away from the benchmark", func() {
				It("returns an empty string, increments the stats for the second bucket", func() {
					s = StatList{
						FirstStatIndex: 0,
						BenchmarkTime:  1000,
						Stats:          [2]Stat{NewStat(stat1SectionHits, stat1Failures), NewStat(stat2SectionHits, stat2Failures)},
					}

					Expect(s.UpdateStatList("some section", 200, 1011)).To(Equal(""))
					Expect(s.BenchmarkTime).To(Equal(int64(1000)))
					Expect(s.FirstStatIndex).To(Equal(0))
					Expect(s.Stats[0].SectionHits).To(Equal(map[string]int{"some section": 1, "some other section": 6}))
					Expect(s.Stats[0].Failures).To(Equal(map[string]int{"some section": 1, "some other section": 1}))
					Expect(s.Stats[1].SectionHits).To(Equal(map[string]int{"some section": 8, "some other section": 5}))
					Expect(s.Stats[1].Failures).To(Equal(map[string]int{"some section": 1, "some other section": 2}))
				})
			})

			Context("The time is 0-10 away from the benchmark", func() {
				It("returns nil and increments the stat for the first bucket", func() {
					s = StatList{
						FirstStatIndex: 0,
						BenchmarkTime:  1000,
						Stats:          [2]Stat{NewStat(stat1SectionHits, stat1Failures), NewStat(stat2SectionHits, stat2Failures)},
					}

					Expect(s.UpdateStatList("some section", 200, 1005)).To(Equal(""))
					Expect(s.BenchmarkTime).To(Equal(int64(1000)))
					Expect(s.FirstStatIndex).To(Equal(0))
					Expect(s.Stats[0].SectionHits).To(Equal(map[string]int{"some section": 2, "some other section": 6}))
					Expect(s.Stats[0].Failures).To(Equal(map[string]int{"some section": 1, "some other section": 1}))
					Expect(s.Stats[1].SectionHits).To(Equal(map[string]int{"some section": 7, "some other section": 5}))
					Expect(s.Stats[1].Failures).To(Equal(map[string]int{"some section": 1, "some other section": 2}))
				})
			})
		})
	})
})
