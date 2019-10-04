package stats_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/athornton2012/http_monitor/stats"
)

var _ = Describe("RollingTraffic", func() {
	var (
		r RollingTrafficList
	)

	Describe("#HandleLog", func() {
		Context("Handling the first log", func() {
			It("increments total traffic and sets the latest date", func() {
				r = NewRollingTrafficList(10, 5)

				Expect(r.HandleLog(12345)).To(Equal(""))
				Expect(r.StartIndex).To(Equal(0))
				Expect(r.LatestDate).To(Equal(int64(12345)))
				Expect(r.TotalTraffic).To(Equal(1))
			})
		})

		Context("When the log date is equal to the latest date", func() {
			Context("When the traffic limit is exceeded", func() {
				Context("When the there is already a high traffic alert in effect", func() {
					It("returns the nil string", func() {
						r = RollingTrafficList{
							TrafficList:      []int{60, 0, 0, 0, 0},
							StartIndex:       0,
							LatestDate:       1234,
							HighTrafficAlert: true,
							Limit:            2,
							TotalTraffic:     60,
							WindowSize:       5,
						}

						Expect(r.HandleLog(1234)).To(Equal(""))
					})
				})

				Context("When there is not high traffic alert currently in effect", func() {
					It("returns a high traffic alert", func() {
						r = RollingTrafficList{
							TrafficList:      []int{60, 0, 0, 0, 0},
							StartIndex:       0,
							LatestDate:       1234,
							HighTrafficAlert: false,
							Limit:            2,
							TotalTraffic:     60,
							WindowSize:       5,
						}

						Expect(r.HandleLog(1234)).To(Equal("HIGH TRAFFIC ALERT\n"))
					})
				})
			})

			Context("When average traffic is under the limit", func() {
				Context("When there a high traffic alert is in effect", func() {
					It("returns an alert saying traffic is backt to normal", func() {
						r = RollingTrafficList{
							TrafficList:      []int{2, 0, 0, 0, 0},
							StartIndex:       0,
							LatestDate:       1234,
							HighTrafficAlert: true,
							Limit:            2,
							TotalTraffic:     2,
							WindowSize:       5,
						}

						Expect(r.HandleLog(1234)).To(Equal("NORMAL TRAFFIC RECOVERED BETWEEN 1233 and 1235\n"))
					})
				})

				Context("When there is not high traffic alert currently in effect", func() {
					It("returns the nil string", func() {
						r = RollingTrafficList{
							TrafficList:      []int{2, 0, 0, 0, 0},
							StartIndex:       0,
							LatestDate:       1234,
							HighTrafficAlert: false,
							Limit:            2,
							TotalTraffic:     2,
							WindowSize:       5,
						}

						Expect(r.HandleLog(1234)).To(Equal(""))
					})
				})
			})
		})

		Context("When the log date is greater than the latest date", func() {
			Context("When the start index is at the end of the traffic list", func() {
				Context("When the average traffic exceeds the limit", func() {
					Context("When there is not a high traffic alert in effect", func() {
						It("returns a high traffic alert", func() {
							r = RollingTrafficList{
								TrafficList:      []int{0, 2, 2, 2, 3},
								StartIndex:       4,
								LatestDate:       1234,
								HighTrafficAlert: false,
								Limit:            2,
								TotalTraffic:     9,
								WindowSize:       5,
							}

							Expect(r.HandleLog(1235)).To(Equal("HIGH TRAFFIC ALERT\n"))
						})
					})
				})

				Context("When the average traffic no longer exceeds the limit", func() {
					Context("When there is a high traffic alert in effect", func() {
						It("returns a message saying traffic has returned to normal", func() {
							r = RollingTrafficList{
								TrafficList:      []int{100, 1, 1, 0, 1},
								StartIndex:       4,
								LatestDate:       1234,
								HighTrafficAlert: true,
								Limit:            2,
								TotalTraffic:     103,
								WindowSize:       5,
							}

							Expect(r.HandleLog(1236)).To(Equal("NORMAL TRAFFIC RECOVERED BETWEEN 1235 and 1237\n"))
						})
					})
				})
			})
		})

		Context("When the log date is less than the latest date", func() {
			Context("When the start index is at the end of the traffict list", func() {
				Context("When the average traffic exceeds the limit", func() {
					Context("When there is not a high traffic alert in effect", func() {
						It("returns a high traffic alert", func() {
							r = RollingTrafficList{
								TrafficList:      []int{6, 0, 0, 1, 2},
								StartIndex:       0,
								LatestDate:       1234,
								HighTrafficAlert: false,
								Limit:            2,
								TotalTraffic:     9,
								WindowSize:       5,
							}

							Expect(r.HandleLog(1232)).To(Equal("HIGH TRAFFIC ALERT\n"))
						})
					})
				})
			})
		})
	})
})
