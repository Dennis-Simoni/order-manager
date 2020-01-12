package order

import (
	"fmt"
	"order-manager/models/item"
	"order-manager/models/order_status"
)

type Order struct {
	ID     string                   `json:"id" `
	Status order_status.OrderStatus `json:"status"`
	Items  []*item.Item             `json:"items"`
}

// Total calculates the total price of an order
func (order Order) Total() float64 {
	var total float64
	for _, v := range order.Items {
		total += v.Price
	}
	return total
}

// PrintOrder formats the order struct into a string representation
func (order Order) PrintOrder() string {
	return fmt.Sprintf("Order ID: %s\n Order Status: %d\n Items: %v\n", order.ID, order.Status, formatItems(order))
}

// ChangeStatus changes the status of an order.
func (order *Order) ChangeStatus(status order_status.OrderStatus) {
	order.Status = status
}

// formatItems is a helper function to PrintOrder() that formats the order items in a string representation.
func formatItems(order Order) string {
	var str string
	for _, v := range order.Items {
		str += fmt.Sprintf("[%s %s %.2f %s]", v.ID, v.Name, v.Price, v.CurrencyCode)
	}
	return str
}
