package main

import (
	"fmt"
	"sort"
	"time"
)

// create a function that takes slice of jobs, find first available slot

type Job struct {
	Start    time.Time
	Duration time.Duration
}

type Jobs []Job

func (j Jobs) Len() int {
	return len(j)
}

func (j Jobs) Less(a, b int) bool {
	return j[a].Start.Before(j[b].Start)
}

func (j Jobs) Swap(a, b int) {
	j[a], j[b] = j[b], j[a]
}

func FindTimeForJob(jobs Jobs, jobDuration time.Duration) time.Time {
	sort.Sort(jobs)

	currentTime := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	for _, j := range jobs {
		availableDuration := j.Start.Sub(currentTime)
		if availableDuration >= jobDuration {
			return currentTime
		}

		currentTime = j.Start.Add(j.Duration)
	}

	return currentTime
}

func main() {
	midnight := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	jobs := Jobs{
		// 12-1
		{midnight, time.Hour},
		// 1-1:30
		{midnight.Add(time.Hour), time.Minute * 30},
		// 2-5
		{midnight.Add(time.Hour * 2), time.Hour * 3},
		// 5-5:15
		{midnight.Add(time.Hour * 5), time.Minute * 15},
		// 6:30-7
		{midnight.Add(time.Hour*6 + time.Minute*30), time.Minute * 30},
	}

	fmt.Println(FindTimeForJob(jobs, time.Hour))
}
