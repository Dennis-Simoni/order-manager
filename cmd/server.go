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
