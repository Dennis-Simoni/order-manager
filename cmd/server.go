package main

import (
	"fmt"
	item "order-manager/models/item"
	order2 "order-manager/models/order"
)

func main() {

	item1 := item.Item{
		ID:           "Test",
		Name:         "123",
		CurrencyCode: "GBP",
		Price:        19.99,
	}

	item2 := item.Item{
		ID:           "Test2",
		Name:         "1234",
		CurrencyCode: "GBP",
		Price:        10.00,
	}

	order := order2.Order{
		ID:     "Order-123",
		Status: "Pending",
		Items:  []*item.Item{&item1, &item2},
	}

	fmt.Println(item1.PrintItem())
	item1.ChangePrice(23, "EUR")

	fmt.Println("Item after changes:\n", item1.PrintItem())

	total := fmt.Sprintf("The order total is: %.2f", order.Total())
	fmt.Println(total)

	fmt.Printf("Printing the order: %v\n", order.PrintOrder())

}