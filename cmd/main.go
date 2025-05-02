package main

import "log"

func main() {
	log.Printf("job starting jobId=%s", jobID)
	if err := run(); err != nil {
		log.Fatalf("job failed jobId=%s: %v", jobID, err)
	}
	log.Printf("job completed jobId=%s", jobID)
}
