package order_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"order-manager/models/item"
	"order-manager/models/order"
	"testing"
)

var item1 = item.Item{
	ID:           "Test",
	Name:         "123",
	CurrencyCode: "GBP",
	Price:        19.99,
}

var item2 = item.Item{
	ID:           "Test2",
	Name:         "1234",
	CurrencyCode: "EUR",
	Price:        10.00,
}

var or = order.Order{
	ID:     "Order-123",
	Status: "Pending",
	Items:  []*item.Item{&item1, &item2},
}

func TestTotal(t *testing.T) {
	expected := 29.99
	got := or.Total()
	assert.Equal(t, got, expected)
}

func TestPrintOrder(t *testing.T) {
	expected := fmt.Sprintf("Order ID: %s\n Order Status: %s\n Items: [%s %s %.2f %s][%s %s %.2f %s]\n",
		or.ID, or.Status, item1.ID, item1.Name, item1.Price, item1.CurrencyCode,
		item2.ID, item2.Name, item2.Price, item2.CurrencyCode)

	got := or.PrintOrder()
	if expected != got {
		t.Errorf("Incorrect result, expected: %v, got: %v", expected, got )
	}
}
