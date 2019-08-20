package main

import (
	"fmt"

	"github.com/zpatrick/etc/ttt/concurrency/shop"
)

var tickets = map[string]string{
	"latte":             "Method Man",
	"cappuccino":        "RZA",
	"macchiato":         "GZA",
	"soy vanilla latte": "Raekwon",
	"espresso":          "Inspectah Deck",
	"americano":         "U-God",
	"mocha":             "ODB",
	"long macchiato":    "Masta Killa",
	"ristretto":         "Ghostface Killah",
	"affogato":          "Capadonna",
}

func main() {
	stand := shop.NewStand()
	for order, customer := range tickets {
		fmt.Printf("Ordering a %s for %s\n", order, customer)
		drink := stand.Order(order, customer)
		fmt.Printf("%s got his %s\n", drink.Customer, drink.Name)
	}
}
