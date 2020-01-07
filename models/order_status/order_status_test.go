package order_status_test

import (
	"order-manager/models/item"
	"order-manager/models/order"
	"order-manager/models/order_status"
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

var newOrder = order.Order{
	ID:     "Order-123",
	Status: order_status.New,
	Items:  []*item.Item{&item1, &item2},
}

var prepOrder = order.Order{
	ID:     "Order-123",
	Status: order_status.Preparing,
	Items:  []*item.Item{&item1, &item2},
}

var readyOrder = order.Order{
	ID:     "Order-123",
	Status: order_status.Ready,
	Items:  []*item.Item{&item1, &item2},
}

var completedOrder = order.Order{
	ID:     "Order-123",
	Status: order_status.Completed,
	Items:  []*item.Item{&item1, &item2},
}

func TestOrderStatus(t *testing.T) {
	if newOrder.Status != 0 {
		t.Errorf("Unexpected status, want: %d, got: %v", order_status.New, newOrder.Status)
	}

	if prepOrder.Status != 1 {
		t.Errorf("Unexpected status, want: %d, got: %v", order_status.Preparing, prepOrder.Status)
	}

	if readyOrder.Status != 2 {
		t.Errorf("Unexpected status, want: %d, got: %v", order_status.Ready, readyOrder.Status)
	}

	if completedOrder.Status != 3 {
		t.Errorf("Unexpected status, want: %d, got: %v", order_status.Completed, completedOrder.Status)
	}
}
