package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"order-manager/models/order"
	"order-manager/repo"
)

type Handler struct {
	Repo *repo.OrderRepo
}

// Index is invoked by HTTP GET /
func (h *Handler) Index() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// Response
		ctx.JSON(http.StatusOK, gin.H{"INFO": "Welcome to the order manager service!"})
	}
}

// GetOrders is invoked by HTTP GET /orders
func (h *Handler) GetOrders() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// Call to the repository method corresponding to the operation
		// Marshal (Serialization) of the return value.
		orders, err := json.Marshal(h.Repo.FetchAll())

		// Error handling, HTTP status & response
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "unable to serve the request"})
			return
		}

		// Perform null check, HTTP status & response
		if string(orders) == "null" {
			ctx.JSON(http.StatusOK, gin.H{"INFO": "currently there are no available orders to display"})
			return
		}
		// Response
		log.Printf("Order string representation: %s", string(orders))
		ctx.JSON(http.StatusOK, gin.H{"INFO": "orders retrieved", "serialized order objects:": orders})
	}
}

// GetOrder is invoked by HTTP GET /orders/{order_id}
func (h *Handler) GetOrder() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// Call the corresponding repository method to fetch the order with
		// the order_id provided on the request body
		or, err := h.Repo.Fetch(ctx.Param("order_id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"Error": error.Error(err)})
			return
		}
		// Serializes the retrieved object
		o, err := json.Marshal(or)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "unable to serve the request"})
			return
		}
		// Response
		log.Printf("Order Retrieved: %v", string(o))
		ctx.JSON(http.StatusOK, o)
	}
}

// PostOrder is invoked by HTTP POST /orders
func (h *Handler) PostOrder() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var order order.Order

		//LimitReader returns a Reader that reads from the request body but stops with EOF after 1 megabyte.
		body, err := ioutil.ReadAll(io.LimitReader(ctx.Request.Body, 1048576))
		// Error handling, HTTP status & response
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "unable to read request body"})
			return
		}
		// Unmarshal request to order struct, HTTP status & response
		if err := json.Unmarshal(body, &order); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"Error": "invalid order body"})
			return
		}
		// Call to the repository method corresponding to the POST operation
		err = h.Repo.DB.UpsertOrder(order)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "the server responded with error"})
			return
		}
		// Response
		ctx.JSON(http.StatusCreated, gin.H{"INFO": "order created"})
	}
}

// DelOrder is invoked by HTTP DELETE /orders/{order_id}
func (h *Handler) DelOrder() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// Call to the repository method corresponding to the DELETE operation
		if err := h.Repo.Delete(ctx.Param("order_id")); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"Error": error.Error(err)})
			return
		}
		// Response
		ctx.JSON(http.StatusOK, gin.H{"INFO": "order deleted"})
	}
}
