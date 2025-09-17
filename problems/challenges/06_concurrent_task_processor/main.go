package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID       int
	Duration time.Duration
}

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d starting job %d\n", id, job.ID)
		time.Sleep(job.Duration)
		fmt.Printf("Worker %d finished job %d\n", id, job.ID)
	}
}

func main() {
	jobList := []Job{
		{ID: 1, Duration: 2 * time.Second},
		{ID: 2, Duration: 1 * time.Second},
		{ID: 3, Duration: 3 * time.Second},
		{ID: 4, Duration: 1 * time.Second},
		{ID: 5, Duration: 2 * time.Second},
	}

	jobsChannel := make(chan Job, len(jobList))

	var wg sync.WaitGroup

	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobsChannel, &wg)
	}

	for _, job := range jobList {
		jobsChannel <- job
	}

	close(jobsChannel)

	wg.Wait()

	fmt.Println("All jobs have been processed.")
}
