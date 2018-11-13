package main

// example program that does jobs
import (
	"log"
	"time"

	"github.com/zpatrick/borg"
)

const (
	addr    = "16.16.16.32"
	timeout = time.Minute
	token   = "some_token"
)

// todo: exchanging tokens?

func main() {
	node, err := borg.Connect(addr, timeout, token)
	if err != nil {
		log.Fatal(err)
	}
	defer node.Disconnect()

	// Initialization: Tells borg how we want to behave.
	// We should add a simple mechanism for users to send messages to the borg,
	// and specify if they are time constrained or not - or to put it another way,
	// we don't want to waste resources sending init messages each time a node joins,
	// if we've gotten this message before, ignore it and don't send it along.
	// anyways, that would be handled by my node's 'Set' helper functions.
	if err := n.SetFoo(foo); err != nil {
		log.Fatal(err)
	}

	// tells node to disconnect if it has 0 connections for 5 minutes
	// notice we are a borg now, anything 'I' do, 'we' do.
	// therefore, we set every node in the borg's lifespan into an hour.
	if err := n.SetLifespan(time.Hour); err != nil {
		log.Fatal(err)
	}

	w := NewWorker(n)
	for e := range n.Events {
		switch e.Data.(type) {
		case borg.Disconnect:
			return
		case borg.Error:
			// switch on err type
			return
		case borg.Job:
			job := e.(borg.Job)
			// totally doing magic here, will change
			jobIDs := job.Data.([]int)
			if err := w.Work(n, jobIDs); err != nil {
				log.Fatal(err)
			}
		}
	}
}
