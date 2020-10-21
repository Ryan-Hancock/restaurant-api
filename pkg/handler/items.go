package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ryan-hancock/resturant-api/pkg/items"
)

type createBurgerRequest struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type changePriceRequest struct {
	Price float32 `json:"price"`
}

type itemResponse struct {
	ItemID int `json:"item_id"`
}

type itemsResponse struct {
	Items []items.Item `json:"items"`
}

// ItemController implementation of ItemService.
type itemController struct {
	s items.Service
}

func newItemController(s items.Service) *itemController {
	return &itemController{
		s: s,
	}
}

func (ic itemController) GetItems(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, itemsResponse{Items: ic.s.GetItems()})

	return
}

func (ic itemController) PostItem(w http.ResponseWriter, r *http.Request) {
	var cr createBurgerRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cr); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid request: %s", err.Error()))
		return
	}
	defer r.Body.Close()

	if cr.Name == "" {
		respondWithError(w, http.StatusBadRequest, validationError{
			Property: "name",
			Message:  "name for item can not be empty",
		})
	}

	id, err := ic.s.NewItem(items.Item{Name: cr.Name, Price: cr.Price})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("unkwone error %s", err.Error()))
	}

	respondWithJSON(w, http.StatusCreated, itemResponse{ItemID: id})

	return
}

func (ic itemController) PatchItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["itemID"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, validationError{
			Property: "itemID",
			Message:  "item ID failed to validate",
		})
		return
	}

	var cr changePriceRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cr); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid request: %s", err.Error()))
		return
	}

	if err = ic.s.ChangeItemPrice(id, cr.Price); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid request: %s", err.Error()))
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)

	return
}
