package main

import "log"

func main() {
	log.Printf("job starting jobId=%s", jobID)
	if err := run(); err != nil {
		log.Fatalf("job failed jobId=%s: %v", jobID, err)
	}
	log.Printf("job completed jobId=%s", jobID)
}

func runSteps(steps []*Step, log *Logger) error {
	for _, step := range steps {
		log.Infof("executing step: %+v", step)
		if err := step.Run(); err != nil {
			log.Errorf("step failed, config: %+v error: %v", step, err)
			return err
		}
		log.Infof("step completed: %+v", step)
	}
	return nil
}
