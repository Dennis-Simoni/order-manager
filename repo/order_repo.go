package repo

import (
	"order-manager/database"
	"order-manager/models/order"
)
// OrderRepo takes the DB struct as a dependency, which represents the database connection.
type OrderRepo struct {
	DB *database.DB
}

// Upsert inserts or updates a model into the database
func (or *OrderRepo) Upsert(order *order.Order) error {
	return or.DB.UpsertOrder(*order)
}

// Fetch returns an order with a specific id from the database
func (or *OrderRepo) Fetch(orderID string) (*order.Order, error) {
	order,err := or.DB.FetchOrder(orderID)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// FetchAll returns all the existing orders from the database
func (or *OrderRepo) FetchAll() []order.Order {
	return or.DB.FetchAllOrders()
}

// Delete deletes a specific order from the database
func (or *OrderRepo) Delete(id string) error {
	return or.DB.DeleteOrder(id)
}