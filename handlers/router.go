package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
)

// Start creates a gin router with default middleware and spins up a server.
func Start(h *Handler) {
	router := gin.Default()
	router.GET("/", h.Index())
	log.Fatal(router.Run(":8080"))
}

/*
func main() {


	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}}
 */