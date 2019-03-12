package main

import (
	"fmt"
	"os"
	"time"
)

type Job struct {
	Start    time.Time
	Duration time.Duration
}

func NewJob(start time.Time, d time.Duration) Job {
	return Job{
		Start: start, 
		Duration: d,
	}
}

func (j Job) End() time.Time {
	return j.Start.Add(j.Duration)
}

func ScheduleJob(jobs []Job, d time.Duration) (time.Time, error) {
	// add a 'job' at the end of the day, so an empty schedule has a 24 hour gap
	jobs = append(jobs, Job{time.Date(0, 0, 0, 24, 0, 0, 0, time.UTC), 0})
	
	current := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	for _, job := range jobs {
		if gap := job.Start.Sub(current); gap >= d {
			return current, nil
		}

		current = job.End().Add(time.Minute)
	}
	
	return time.Time{}, fmt.Errorf("Cannot schedule a %s job", d)
}

func main() {
	jobs := []Job{
		// 5 to 30 (00:05 to 00:30)
		NewJob(time.Date(0, 0, 0, 0, 5, 0, 0, time.UTC), time.Minute*25),
		// 120 to 241 (02:00 to 04:01)
		NewJob(time.Date(0, 0, 0, 2, 0, 0, 0, time.UTC), time.Hour*2+time.Minute),
		// 790 to 1015 (13:10 to 16:55)
		NewJob(time.Date(0, 0, 0, 13, 10, 0, 0, time.UTC), time.Hour*3+time.Minute*45),
	}

	start, err := ScheduleJob(jobs, time.Hour*3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Job can be scheduled at:", start)
}
