package shop

import (
	"fmt"
)

type order struct {
	drink string
	customer string
}

type BrickAndMortar struct {
}

func NewBrickAndMortar() *BrickAndMortar {
	return &BrickAndMortar{}
}

// also helsp with online orders
func (b *BrickAndMortar) Open() func() {
	orders := make(chan Order)
	go func() {
		for order := range orders {
			
		}
	}()

	return func() { close(orders) }
}

func (s *BrickAndMortar) Order(drink, customer string) Drink {
	s.orders <- Order{drink: drink, customer: customer}

	fmt.Printf("Brewing %s for %s\n", drink, customer)
	return Drink{Name: drink, Customer: customer}
}
