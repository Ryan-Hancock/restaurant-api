package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ryan-hancock/resturant-api/pkg/items"
)

type createBurgerRequest struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type itemResponse struct {
	ItemID int `json:"item_id"`
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
	return
}

func (ic itemController) PostItem(w http.ResponseWriter, r *http.Request) {
	var cr createBurgerRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cr); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid request: %s", err))
		return
	}
	defer r.Body.Close()

	ID, err := ic.s.NewItem(items.Item{Name: cr.Name, Price: cr.Price})
	if errors.Is(err, items.ErrNotFound) {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("unkwone error %s", err))
	}

	respondWithJSON(w, http.StatusCreated, itemResponse{ItemID: ID})

	return
}

func (ic itemController) PatchItem(w http.ResponseWriter, r *http.Request) {
	return
}
