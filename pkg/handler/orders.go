package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ryan-hancock/resturant-api/pkg/items"
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
	vars := mux.Vars(r)

	orderID, err := strconv.Atoi(vars["orderID"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, validationError{
			Property: "orderID",
			Message:  "order ID failed to validate",
		})
	}

	itemID, err := strconv.Atoi(vars["itemID"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, validationError{
			Property: "itemID",
			Message:  "item ID failed to validate",
		})
	}

	_, err = oc.s.AppendLine(orders.Line{OrderID: orderID, ItemID: itemID})
	if err == orders.ErrNotFound {
		respondWithError(w, http.StatusNotFound, validationError{
			Property: "orderID",
			Message:  "could not find order",
		})
		return
	} else if err == items.ErrNotFound {
		respondWithError(w, http.StatusNotFound, validationError{
			Property: "itemID",
			Message:  "could not find item",
		})
		return
	}
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid request: %s", err))
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)

	return
}

func (oc orderController) PostOrderPay(w http.ResponseWriter, r *http.Request) {
	return
}
