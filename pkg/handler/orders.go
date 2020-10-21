package handler

import (
	"fmt"
	"net/http"

	"github.com/ryan-hancock/resturant-api/pkg/orders"
)

type orderController struct {
	s orders.Service
}

func newOrderController(s orders.Service) *orderController {
	return &orderController{
		s: s,
	}
}

type orderResponse struct {
	OrderID int `json:"order_id"`
}

func (oc orderController) PostOrder(w http.ResponseWriter, r *http.Request) {
	id, err := oc.s.NewOrder()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid request: %s", err))
	}

	respondWithJSON(w, http.StatusCreated, orderResponse{OrderID: id})

	return
}

func (oc orderController) PatchOrderWithItem(w http.ResponseWriter, r *http.Request) {
	return
}

func (oc orderController) PostOrderPay(w http.ResponseWriter, r *http.Request) {
	return
}
