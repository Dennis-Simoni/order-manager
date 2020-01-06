package main

import (
	"fmt"
	"order-manager/models"
)

func main() {
	item := models.Item{
		ID:           "Test",
		Name:         "123",
		CurrencyCode: "GBP",
		Price:        19.99,
	}

	fmt.Println(item.PrintItem())
	item.ChangePrice(23, "EUR")

	fmt.Println("Item after changes:\n", item.PrintItem())
}