package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"order-manager/repo"
)

type Handler struct {
	Repo *repo.OrderRepo
}

func (h *Handler) Index() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to the Order Manager Service!")
	}
}

func (h *Handler) GetOrders() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		orders, err := json.Marshal(h.Repo.FetchAll())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error":"Unable to serve the request"})
			return
		}

		if string(orders) == "null" {
			ctx.JSON(http.StatusOK, "Currently there are no available orders to display!")
			return
		}
		ctx.JSON(http.StatusOK, orders)
	}
}

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) PutOrder(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) DelOrder(w http.ResponseWriter, r *http.Request)  {}
