package main

import (
	"fmt"
	"time"
)

type Job struct {
	ID int
}

type Pool struct {
	workers []*Worker
}

func NewPool(numWorkers int) *Pool {
	workers := make([]*Worker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = NewWorker(i)
	}

	return &Pool{
		workers: workers,
	}
}

func (p *Pool) Listen(c <-chan Job) {
	for _, w := range p.workers {
		go w.Listen(c)
	}
}

type Worker struct {
	ID int
}

func NewWorker(id int) *Worker {
	return &Worker{ID: id}
}

func (w *Worker) Listen(c <-chan Job) {
	for j := range c {
		fmt.Printf("[Worker %d] Received Job %d\n", w.ID, j.ID)
		time.Sleep(time.Second)
		fmt.Printf("[Worker %d] Job %d completed\n", w.ID, j.ID)
	}
	fmt.Printf("[Worker %d] Done working\n", w.ID)
}

func main() {
	c := make(chan Job)
	pool := NewPool(2)
	pool.Listen(c)

	for i := 0; i < 5; i++ {
		c <- Job{ID: i}
	}
	close(c)
	time.Sleep(time.Second * 2)
}
