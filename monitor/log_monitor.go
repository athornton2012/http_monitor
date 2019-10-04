package monitor

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/athornton2012/http_monitor/stats"
)

const HeaderLog = "\"remotehost\",\"rfc931\",\"authuser\",\"date\",\"request\",\"status\",\"bytes\""

type LogMonitor struct {
	Writer             *bufio.Writer
	StatList           stats.StatList
	RollingTrafficList stats.RollingTrafficList
	LogStream          chan string
	Done               chan bool
}

func NewLogMonitor(writer *bufio.Writer, logStream chan string, done chan bool, alertLimit int) LogMonitor {
	return LogMonitor{
		Writer:             writer,
		StatList:           stats.NewStatList(),
		RollingTrafficList: stats.NewRollingTrafficList(alertLimit, 120),
		LogStream:          logStream,
		Done:               done,
	}
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . StatList
type StatList interface {
	UpdateStatList() string
	FlushAll() string
}

//counterfeiter:generate . RollingTrafficList
type RollingTrafficList interface {
	HandleLog(date int64) string
}

type LogLine struct {
	RemoteHost string
	Rfc93      string
	AuthUser   string
	Date       int64
	Request    string
	Status     string //todo
	Bytes      string
}

func NewLogLine(remoteHost, rfc93, authUser string, date int64, request, status, bytes string) LogLine {
	return LogLine{
		RemoteHost: remoteHost,
		Rfc93:      rfc93,
		AuthUser:   authUser,
		Date:       date,
		Request:    request,
		Status:     status,
		Bytes:      bytes,
	}
}

func TokensToLogLine(logTokens []string) LogLine {
	logTime, err := strconv.ParseInt(logTokens[4], 10, 64)
	if err != nil {
		fmt.Println("Unable to parse timestamp:", err)
	}

	return NewLogLine(
		logTokens[1],
		logTokens[2],
		logTokens[3],
		logTime,
		logTokens[5],
		logTokens[6],
		logTokens[7],
	)
}

func (lm LogMonitor) Monitor() {
	for {
		log, more := <-lm.LogStream
		if more {
			if log == HeaderLog {
				continue
			}

			tokenDelimiters := regexp.MustCompile(`["]*,["]*|"`)
			logTokens := tokenDelimiters.Split(log, -1)

			currentLine := TokensToLogLine(logTokens)
			endpoint := strings.Split(currentLine.Request, " ")[1]
			section := strings.Split(endpoint, "/")[1]
			code, err := strconv.Atoi(currentLine.Status)
			if err != nil {
				io.WriteString(lm.Writer, "Unable to convert status code to int:"+err.Error())
			}

			trafficStats := lm.StatList.UpdateStatList(section, code, currentLine.Date)
			alert := lm.RollingTrafficList.HandleLog(currentLine.Date)

			if trafficStats != "" {
				io.WriteString(lm.Writer, trafficStats)
			}

			if alert != "" {
				io.WriteString(lm.Writer, alert)
			}

		} else {
			io.WriteString(lm.Writer, lm.StatList.FlushAll())
			lm.Done <- true
			io.WriteString(lm.Writer, "Finished")
			return
		}
	}
}
