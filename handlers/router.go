package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
)

// Start creates a gin router with default middleware and spins up a server.
func Start(h *Handler) {
	router := gin.Default()

	router.GET("/", h.Index())
	router.GET("/orders", h.GetOrders())

	log.Fatal(router.Run(":8080"))
}

// go test ./... to run all tests