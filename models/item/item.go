package item

import "fmt"

type Item struct {
	ID, Name, CurrencyCode string
	Price                  float64
}

// PrintItem prints the Item struct representation in a form of string.
func (i Item) PrintItem() string {
	return fmt.Sprintf("ID: %s\n"+
		"Name: %s\nCurrency Code: %s\nPrice: %.2f\n", i.ID, i.Name, i.CurrencyCode, i.Price)
}

func (i *Item) ChangeID(id string) {
	i.ID = id
}

func (i *Item) ChangeName(name string) {
	i.Name = name
}

func (i *Item) ChangePrice(price float64, code string) {
	i.Price = price
	i.CurrencyCode = code
}