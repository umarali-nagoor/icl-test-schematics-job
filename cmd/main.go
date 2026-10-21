package main

import "log"

func main() {
	log.Printf("job starting jobId=%s", jobID)
	if err := run(); err != nil {
		log.Fatalf("job failed jobId=%s: %v", jobID, err)
	}
	log.Printf("job completed jobId=%s", jobID)
}

func watchJobStatus(jobID string, client JobClient, log *Logger) {
	for {
		job, err := client.Get(jobID)
		if err != nil {
			log.Errorf("status poll failed for job %s: %v", jobID, err)
		} else {
			log.Infof("job status: %+v", job)
		}
		time.Sleep(2 * time.Second)
	}
}
