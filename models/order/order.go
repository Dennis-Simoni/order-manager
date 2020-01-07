package order

import (
	"fmt"
	"order-manager/models/item"
)

type Order struct {
	ID, Status string
	Items      []*item.Item
}

func (order Order) Total() float64 {
	var total float64
	for _, v := range order.Items {
		total += v.Price
	}
	return total
}

func (order Order) PrintOrder() string {
	return fmt.Sprintf("Order ID: %s\n Order Status: %s\n Items: %v\n", order.ID, order.Status, formatItems(order))
}

func formatItems(order Order) string {
	var str string
	for _, v := range order.Items {
		str += fmt.Sprintf("[%s %s %.2f %s]", v.ID, v.Name, v.Price, v.CurrencyCode)
	}
	return str
}
