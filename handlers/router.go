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
	router.GET("/orders/:order_id", h.GetOrder())
	router.POST("/orders", h.PostOrder())
	router.DELETE("/orders/:order_id", h.DelOrder())

	log.Fatal(router.Run(":8080"))
}
