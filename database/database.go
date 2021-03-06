package database

import (
	"errors"
	"fmt"
	"order-manager/models/order"
	"sync"
)

// DB is a struct that defines an in-memory storage.
type DB struct {
	m sync.Mutex
	orders map[string]order.Order
}

// InitDB is an initializer function for the database.
func InitDB() *DB {
	return &DB{
		orders: make(map[string]order.Order),
	}
}
// UpsertOrder order allows to atomically insert a record
func (db *DB) UpsertOrder(order order.Order) error {
	db.m.Lock()
	defer db.m.Unlock()
	if _, ok := db.orders[order.ID]; !ok {
		order.ID = fmt.Sprintf("order-%d", len(db.orders))
	}
	db.orders[order.ID] = order
	return nil
}

// FetchOrder receives an order ID and if a record exists returns it, otherwise returns an error.
func (db *DB) FetchOrder(orderID string) (order.Order, error) {
	or, ok := db.orders[orderID]
	if !ok {
		return order.Order{}, errors.New(fmt.Sprintf("The order id: %s is not found", orderID))
	}
	return or, nil
}

// FetchAllOrders returns all records available, if any.
func (db *DB) FetchAllOrders() []order.Order {
	var orders []order.Order

	for _, v := range db.orders {
		orders = append(orders, v)
	}

	return orders
}

// DeleteOrder receives an order ID and if a record exists, deletes it, otherwise returns an error.
func (db *DB) DeleteOrder(orderID string) error {
	db.m.Lock()
	defer db.m.Unlock()
	_, ok := db.orders[orderID]
	if !ok {
		return errors.New(fmt.Sprintf("The order id %s doesn't match an active record", orderID))
	}

	delete(db.orders, orderID)
	return nil
}