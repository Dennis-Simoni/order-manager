package order_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"order-manager/models/item"
	"order-manager/models/order"
	"order-manager/models/order_status"
	"testing"
)

const newStatus = order_status.Cancelled

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
	Status: order_status.Preparing,
	Items:  []*item.Item{&item1, &item2},
}

func TestTotal(t *testing.T) {
	expected := 29.99
	got := or.Total()
	assert.Equal(t, got, expected)
}

func TestPrintOrder(t *testing.T) {
	want := fmt.Sprintf("Order ID: %s\n Order Status: %d\n Items: [%s %s %.2f %s][%s %s %.2f %s]\n",
		or.ID, or.Status, item1.ID, item1.Name, item1.Price, item1.CurrencyCode,
		item2.ID, item2.Name, item2.Price, item2.CurrencyCode)

	got := or.PrintOrder()
	if want != got {
		t.Errorf("Incorrect result, got: %v, want: %v", got, want)
	}
}

func TestChangeStatus(t *testing.T) {
	or.ChangeStatus(newStatus)
	if or.Status != newStatus {
		t.Errorf("Order status incorrect, got: %d want: %v", order_status.Cancelled, newStatus)
	}
}

func BenchmarkChangeStatus(b *testing.B) {
	// run the Change Name function b.N times
	for n := 0; n < b.N; n++ {
		or.ChangeStatus(3)
	}
}