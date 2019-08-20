package shop

type Shop interface {
	Order(drink, customer string) (Drink, error)
}
