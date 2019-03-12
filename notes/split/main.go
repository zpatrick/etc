package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

type Job struct {
	ID string
}

type JobRequest struct {
	Job  Job
	Errc chan<- error
}

type Worker struct {
	ID string
}

func (w Worker) Listen(jrc <-chan JobRequest) {
	for jr := range jrc {
		w.Perform(jr.Job, jr.Errc)
	}
}

func (w *Worker) Perform(j Job, errc chan<- error) {
	fmt.Printf("[Worker %v] Performing job %v\n", w.ID, j.ID)
	errc <- nil
}

func Connect(workers []*Worker) (sender func(j Job) error, closer func()) {
	jobc := make(chan JobRequest)
	for _, worker := range workers {
		go worker.Listen(jobc)
	}

	sender = func(j Job) error {
		errc := make(chan error)
		defer close(errc)

		jobc <- JobRequest{j, errc}
		return <-errc
	}

	closer = func() {
		close(jobc)
	}

	return sender, closer
}

func main() {
	numWorkers := flag.Int("workers", 2, "the number of workers")
	numJobs := flag.Int("jobs", 50, "the number of jobs")
	flag.Parse()

	workers := make([]*Worker, *numWorkers)
	for i := range workers {
		workers[i] = &Worker{ID: strconv.Itoa(i)}
	}

	send, close := Connect(workers)
	defer close()

	jobs := make([]Job, *numJobs)
	for i := range jobs {
		jobs[i] = Job{ID: strconv.Itoa(i)}
	}

	fmt.Printf("Spreading %d jobs across %d workers\n", *numJobs, *numWorkers)
	fmt.Printf("Press the Enter Key to continue...\n")
	fmt.Scanln()

	start := time.Now()
	// todo: async publish jobs; like jobs coming in from a web request
	// todo: build httptest server and publish results to myself here; mock a whole bunch of web requests
	for _, job := range jobs {
		go func() {
			if err := send(job); err != nil {
				fmt.Printf("ERROR: [Job %v] %v\n", job.ID, err)
			}
		}()
	}

	fmt.Printf("Done (%s)\n", time.Since(start))
}
