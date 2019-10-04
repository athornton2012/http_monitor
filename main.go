package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/athornton2012/http_monitor/monitor"
)

func main() {
	logFilePath := flag.String("logFilePath", "", "full path to log file")
	alertThreshold := flag.Int("alertThreshold", 10, "the average amount of requests per second that trigger an alert")

	flag.Parse()

	logFileReader, err := os.Open(*logFilePath)
	if err != nil {
		fmt.Println("Unable to open log file", err)
		os.Exit(1)
	}

	requests := make(chan string, 100)
	done := make(chan bool)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()

	go func() {
		logMonitor := monitor.NewLogMonitor(f, requests, done, *alertThreshold)
		logMonitor.Monitor()
	}()

	scanner := bufio.NewScanner(logFileReader)
	for scanner.Scan() {
		requests <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading logFile:", err)
		os.Exit(1)
	}

	close(requests) //done scanning request file line by line

	<-done
}
