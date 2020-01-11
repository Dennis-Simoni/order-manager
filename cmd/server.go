package main

import (
	"order-manager/database"
	"order-manager/handlers"
	"order-manager/repo"
)

func main() {

	orderRepo := &repo.OrderRepo{
		DB: database.InitDB(),
	}

	handler := &handlers.Handler{
		Repo: orderRepo,
	}

	handlers.Start(handler)
}

/*
Welcome/Index page
	HTTP GET /
Find all orders
	HTTP GET /orders
Find order with given id
	HTTP GET /orders/:orderID
Upsert order
	HTTP POST /orders
Delete order with given id
	HTTP DELETE /orders/id
*/