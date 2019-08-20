package worker

import "github.com/getoutreach/logutil"

/*
Design:
	We have a first-class concept of jobs: job.Type, job.Status, etc.
	Jobs are submitted via a job.Submitter.
	During the job.Submitter, jobs are validated.
	When a valid job is submitted, it goes into Redis.

	- or -
	the worker pool operates totally indepent based on GUIDs


	really, we need a mechanism to handle:
	- pull jobs from the queue
	- synchronize workers
	- backoff based on queue
*/

type Pool struct {
	logger logutil.Logger
}

func NewPool(logger logutil.Logger) *Pool {
	return &Pool{
		logger: logger,
	}
}

func (p *Pool) Listen() (submit func(jobID string), closer func()) {
	jobc := make(chan string)
	go func() {
		for jobID := range jobc {
			logutil.Infof(p.logger, "Got job %s", jobID)
		}
	}()

	return func(jobID string) { jobc <- jobID }, func() { close(jobc) }
}
