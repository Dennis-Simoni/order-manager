package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-manager/repo"
)

type Handler struct {
	Repo *repo.OrderRepo
}

func (h *Handler) Index() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to the Order Manager Service!")
	}
}
func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) PutOrder(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) DelOrder(w http.ResponseWriter, r *http.Request)  {}
