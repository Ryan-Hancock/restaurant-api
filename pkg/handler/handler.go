package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ryan-hancock/resturant-api/pkg/items"
	"github.com/ryan-hancock/resturant-api/pkg/orders"
)

// Setup will initalise the app handlers for the router.
func Setup(router *mux.Router, ir items.Repository, or orders.Repository) {
	is := items.NewService(ir)
	ic := newItemController(is)

	router.HandleFunc("/item", ic.GetItems).Methods("GET")
	router.HandleFunc("/item", ic.PostItem).Methods("POST")
	router.HandleFunc("/item/{itemID}", ic.PatchItem).Methods("PATCH")

	os := orders.NewService(or)
	oc := newOrderController(os)

	router.HandleFunc("/order", oc.PostOrder).Methods("POST")
	router.HandleFunc("/order/{orderID}/additem/{itemID}", oc.PatchOrderWithItem).Methods("PATCH")
	router.HandleFunc("/order/{orderID}/pay", oc.PostOrderPay).Methods("POST")
}

type errorResponse struct {
	Error interface{} `json:"error"`
}

type validationError struct {
	Property string `json:"property"`
	Message  string `json:"message"`
}

func respondWithError(w http.ResponseWriter, code int, payload interface{}) {
	respondWithJSON(w, code, errorResponse{Error: payload})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
