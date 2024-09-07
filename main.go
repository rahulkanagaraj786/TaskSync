package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Time window to process jobs
var maxProcessingWindow int64 = 10

// Buffered channel to queue messages for processing
var jobQueue = make(chan string, 100)

// Structure representing the job request
type JobRequest struct {
	JobID          string
	ScheduleTime   int64
	Interval       int64
	TargetIP       string
	TargetPort     string
	TargetExchange string
}

func main() {
	defer redisClient.Close() // Ensure Redis client is closed gracefully

	// Launch routine to monitor missed jobs
	go missedJobHandler()

	// Start 5 worker routines to handle job processing
	for i := 0; i < 5; i++ {
		go processJob(i)
	}

	// Start 5 routines to send messages to MQ
	for i := 0; i < 5; i++ {
		go sendToMQ()
	}

	// Uncomment the test function for scheduling sample jobs
	// runTests()

	// Set up HTTP server to handle job creation requests
	http.HandleFunc("/", handleJobCreation)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// Function to simulate test cases
func runTests() {
	// Schedule a large number of jobs simultaneously for testing
	for i := 0; i < 1000000; i++ {
		scheduleJob(float64(currentTimestamp()+5), fmt.Sprintf("%d", i))
	}
}

// Get the current timestamp in seconds since epoch
func currentTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}
