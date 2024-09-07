package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTP handler for job creation
func handleCreateJob(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Println("Received job request:", string(data))

	var jobRequest Request
	if err := json.Unmarshal(data, &jobRequest); err != nil {
		log.Println("Error while parsing JSON:", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validate and schedule job
	log.Printf("Scheduling Job: %s, Scheduled Time: %f", jobRequest.JobName, jobRequest.RunTime)
	addJob(jobRequest.RunTime, jobRequest.JobName)
}
