package main

import (
	"fmt"
	"log"
	"order-manager/database"
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

	order2 := order.Order{
		ID:     "",
		Status: 0,
		Items:  []*item.Item{&item1, &item2},
	}


	fmt.Println(item1.PrintItem())
	item1.ChangePrice(23, "EUR")

	fmt.Println("Item after changes:\n", item1.PrintItem())

	total := fmt.Sprintf("The order total is: %.2f", order1.Total())
	fmt.Println(total)

	fmt.Printf("Printing the order: %v\n", order1.PrintOrder())

	// Store and retrieve values from the in-memory database:
	db := database.InitDB()
	err := db.UpsertOrder(order2)
	if err != nil {
		log.Fatal("Unable to store the order.")
	}

	findOrder(db)

	// find all orders:
	orders, err := db.FetchAllOrders()
	if err != nil {
		log.Fatal("Unable to retrieve the orders.")
	}

	for i, v := range orders {
		fmt.Printf("Order: %d\n",i)
		fmt.Printf("\nOrder ID: %s\n Order Status: %d\n Order Items %q\n", v.ID, v.Status, &v.Items)
	}

	// delete an order:
	err = db.DeleteOrder("order-0")
	if err != nil {
		log.Fatal("Unable to delete the record.")
	}
	// expected to return an error
	findOrder(db)
}

func findOrder(db *database.DB)  {
	o, err := db.FetchOrder("order-0")
	if err != nil {
		log.Fatal("Unable to retrieve the order.")
	}

	f := fmt.Sprintf("\nOrder ID: %s\n Order Status: %d\n Order Items %q\n", o.ID, o.Status, &o.Items)
	log.Println(f)
}