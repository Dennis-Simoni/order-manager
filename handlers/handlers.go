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

func (h *Handler) Index() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"INFO:":"Welcome to the order manager service!"})
	}
}

func (h *Handler) GetOrders() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// Call to the repository method corresponding to the operation
		// Marshal (Serialization) of the return value.
		orders, err := json.Marshal(h.Repo.FetchAll())

		// Error handling, HTTP status & response
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error":"unable to serve the request"})
			return
		}

		// Perform null check, HTTP status & response
		if string(orders) == "null" {
			ctx.JSON(http.StatusOK, gin.H{"INFO:":"currently there are no available orders to display"})
			return
		}
		log.Printf("Order string representation: %s", string(orders))
		ctx.JSON(http.StatusOK, gin.H{"INFO:": "orders retrieved", "serialized order objects:": orders})
	}
}

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {}

// run server: go run server.go
// post request: curl -v -H "Content-Type: application/json"  --data @postBody.json http://localhost:8080/orders
func (h *Handler) PostOrder() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var order order.Order

		//LimitReader returns a Reader that reads from the request body but stops with EOF after 1 megabyte.
		body, err := ioutil.ReadAll(io.LimitReader(ctx.Request.Body, 1048576))
		// Error handling, HTTP status & response
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error:":"unable to read request body"})
		}
		// Unmarshal request to order struct, HTTP status & response
		if err := json.Unmarshal(body, &order); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"Error": "invalid order body"})
		}
		// Call to the repository method corresponding to the POST operation
		err = h.Repo.DB.UpsertOrder(order)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error:":"the server responded with error",})
		}
		// HTTP success status & the value return from the repo
		ctx.JSON(http.StatusCreated, gin.H{"INFO:": "order created"})
	}
}

func (h *Handler) DelOrder(w http.ResponseWriter, r *http.Request) {}