package handler

import (
	"net/http"

	"github.com/ryan-hancock/resturant-api/pkg/orders"
)

type orderController struct {
	service orders.Service
}

func newOrderController(s orders.Service) *orderController {
	return &orderController{
		service: s,
	}
}

func (oc orderController) PostOrder(w http.ResponseWriter, r *http.Request) {
	return
}

func (oc orderController) PatchOrderWithItem(w http.ResponseWriter, r *http.Request) {
	return
}

func (oc orderController) PostOrderPay(w http.ResponseWriter, r *http.Request) {
	return
}
