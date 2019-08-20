package shop

import (
	"fmt"
	"sync"
)

type Stand struct {
	lock sync.Mutex
}

func NewStand() *Stand {
	return &Stand{
		lock: sync.Mutex{},
	}
}

func (s *Stand) Order(drink, customer string) Drink {
	s.lock.Lock()
	defer s.lock.Unlock()

	fmt.Printf("Brewing %s for %s\n", drink, customer)
	return Drink{Name: drink, Customer: customer}
}
