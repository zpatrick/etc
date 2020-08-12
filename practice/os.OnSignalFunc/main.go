package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

type Service struct {
	name string
}

func NewService(name string) *Service {
	return &Service{name}
}

func (s *Service) Start(ctx context.Context, d time.Duration) {
	go func() {
		fmt.Printf("Starting service %s\n", s.name)
		for {
			select {
			case <-time.After(d):
				s.Run()
			case <-ctx.Done():
				fmt.Printf("Stopping service %s\n", s.name)
				return
			}
		}
	}()
}

func (s *Service) Run() error {
	fmt.Printf("Running service %s\n", s.name)
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	NewService("alfa").Start(ctx, time.Second)
	NewService("bravo").Start(ctx, time.Second*2)
	NewService("charlie").Start(ctx, time.Second*3)

	// should be os.OnSignalFunc
	OnSignalFunc(os.Interrupt, func() {
		cancel()
	})

	time.Sleep(time.Second * 5)
	cancel()
	fmt.Println("Done")
}

// here's the new function to add as a helper:to the os package
func OnSignalFunc(s os.Signal, f func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, s)
	go func() {
		<-c
		f()
	}()
}
