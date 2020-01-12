package item

import "fmt"

type Item struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	CurrencyCode string  `json:"currency_code"`
	Price        float64 `json:"price"`
}

// PrintItem prints the Item struct representation in a form of string.
func (i Item) PrintItem() string {
	return fmt.Sprintf("ID: %s\n"+
		"Name: %s\nCurrency Code: %s\nPrice: %.2f\n", i.ID, i.Name, i.CurrencyCode, i.Price)
}

// ChangeID changes the item id.
func (i *Item) ChangeID(id string) {
	i.ID = id
}

// ChangeName changes the item name.
func (i *Item) ChangeName(name string) {
	i.Name = name
}

// ChangePrice changes the item price and currency code.
func (i *Item) ChangePrice(price float64, code string) {
	i.Price = price
	i.CurrencyCode = code
}
