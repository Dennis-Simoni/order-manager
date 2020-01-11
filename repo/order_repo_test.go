package repo_test

import (
	"github.com/stretchr/testify/assert"
	"order-manager/database"
	"order-manager/models/order"
	"order-manager/models/order_status"
	"order-manager/repo"
	"testing"
)

var (
	db = database.InitDB()
	or = &repo.OrderRepo{DB: db}
	_  = or.Upsert(&order.Order{Status: order_status.New})
	_  = or.Upsert(&order.Order{Status: order_status.Ready})
	o3 = or.Upsert(&order.Order{Status: order_status.Preparing})
)

func TestUpsert(t *testing.T) {
	err := o3
	if err != nil {
		t.Fatalf("error should not have been returned from Upsert, got: %v", err)
	}
	got, err := or.Fetch("order-2")
	if got.Status != order_status.Preparing {
		t.Fatalf("incorrect status persisted to DB, got: %v, want: %v", got.Status, order_status.Preparing)
	}
	if err != nil {
		t.Fatalf("model should have been found :%v", err)
	}
}

func TestFetch(t *testing.T) {
	existingOrder, err := or.Fetch("order-0")
	if existingOrder.Status != order_status.New {
		t.Fatalf("incorrect order status  returned got: %v, want: %v", existingOrder.Status, order_status.New)
	}
	if err != nil {
		t.Fatalf("repo should not have returned error, got: %v", err)
	}
	noOrder, err := or.Fetch("model-10")
	if noOrder != nil {
		t.Fatalf("no order should have been found, got: %v", noOrder)
	}
	if err == nil {
		t.Fatalf("error should not be nil when no error exists")
	}
}

func TestFetchAll(t *testing.T) {
	got := or.FetchAll()
	if got[0].Status != order_status.New {
		t.Fatalf("status of first order is incorrect, got: %d, want: %d", got[0].Status, order_status.New)
	}
	if got[1].Status != order_status.Ready {
		t.Fatalf("status of second order is incorrect, got: %d, want: %d", got[1].Status, order_status.Ready)
	}
}

func TestDelete(t *testing.T) {
	if err := or.Delete("order-0"); err != nil {
		t.Fatalf("delete operation should not have returned error, got : %v", err)
	}
	err := or.Delete("order-0")
	if err == nil {
		t.Fatalf("delete operation should have returned error when no key to be deleted")
	}
	orders := or.FetchAll()
	assert.Equal(t, 2, len(orders))
}

func BenchmarkOrderRepo_FetchAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		or.FetchAll()
	}
}
