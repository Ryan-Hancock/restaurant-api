package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/ryan-hancock/restaurant-api/pkg/items"
	"github.com/ryan-hancock/restaurant-api/pkg/storage/memory"
	"github.com/ryan-hancock/restaurant-api/pkg/test"
)

func newTestItemHandler() *itemController {
	ir := memory.NewItemRepository()
	is := items.NewService(ir)
	return newItemController(is)
}

func TestGetItems(t *testing.T) {
	ic := newTestItemHandler()
	req := test.NewRequest(t, "GET", "/item", nil)
	rr := test.ServeRequest("/item", ic.GetItems, req)

	t.Run("I should see a 200 response", func(t *testing.T) {
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

	return
}

func TestPostItem(t *testing.T) {
	ic := newTestItemHandler()

	var jsonStr = []byte(`{"name":"burger", "price": 1.99}`)
	req := test.NewRequest(t, "POST", "/item", bytes.NewBuffer(jsonStr))
	rr := test.ServeRequest("/item", ic.PostItem, req)

	t.Run("I should see a 201 response", func(t *testing.T) {
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}
	})

	jsonStr = []byte(`{"price": 1.99}`)
	req = test.NewRequest(t, "POST", "/item", bytes.NewBuffer(jsonStr))
	rr = test.ServeRequest("/item", ic.PostItem, req)

	t.Run("If I send a request without a name I should get a 400", func(t *testing.T) {
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})

	return
}

func TestPatchItem(t *testing.T) {
	ic := newTestItemHandler()
	iid, _ := ic.s.NewItem(items.Item{Name: "burger", Price: 5.99})

	var jsonStr = []byte(`{"price": 1.99}`)
	req := test.NewRequest(t, "POST", fmt.Sprintf("/item/%d", iid), bytes.NewBuffer(jsonStr))
	rr := test.ServeRequest("/item/{itemID}", ic.PatchItem, req)

	t.Run("I should see a 204 response", func(t *testing.T) {
		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNoContent)
		}
	})

	t.Run("The price should be 1.99", func(t *testing.T) {
		if item, err := ic.s.GetItem(iid); err != nil || item.Price != 1.99 {
			t.Errorf("item returned wrong price: got %v want %v",
				item.Price, 1.99)
		}
	})

	return
}
