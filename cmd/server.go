package main

import (
	"fmt"
	"order-manager/models/item"
	"order-manager/models/order"
	"order-manager/models/order_status"
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

	order1 := order.Order{
		ID:     "Order-123",
		Status: order_status.New,
		Items:  []*item.Item{&item1, &item2},
	}

	fmt.Println(item1.PrintItem())
	item1.ChangePrice(23, "EUR")

	fmt.Println("Item after changes:\n", item1.PrintItem())

	total := fmt.Sprintf("The order total is: %.2f", order1.Total())
	fmt.Println(total)

	fmt.Printf("Printing the order: %v\n", order1.PrintOrder())

}