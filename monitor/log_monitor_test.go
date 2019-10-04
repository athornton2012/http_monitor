package monitor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/athornton2012/http_monitor/monitor"

	"github.com/athornton2012/http_monitor/monitor/monitorfakes"
)

var _ = Describe("LogMonitor", func() {
	var (
		logStrings = []string{
			"'remotehost','rfc93','authuser','date','request','status','bytes'",
			"'10.0.0.2','-','apache',1549573860,'GET /api/user HTTP/1.0',200,1234",
			"'10.0.0.4','-','apache',1549573860,'GET /api/user HTTP/1.0',200,1234",
			"'10.0.0.4','-','apache',1549573860,'GET /api/user HTTP/1.0',200,1234",
		}
		logMonitor LogMonitor
		logStream  chan string
		done       chan bool
	)

	BeforeEach(func() {
		logStream = make(chan string, 10)
		done = make(chan bool)

		logMonitor = LogMonitor{
			Writer.
				LogStream: logStream,
			Done: done,
		}
	})

	Describe("Monitor", func() {
		It("writes ")
	})
})
