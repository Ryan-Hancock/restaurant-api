package handler

import (
	"net/http"
	"strings"
	"testing"

	"github.com/ryan-hancock/resturant-api/pkg/orders"
	"github.com/ryan-hancock/resturant-api/pkg/storage/memory"
	"github.com/ryan-hancock/resturant-api/pkg/test"
)

func newTestOrdersHandler() *orderController {
	or := memory.NewOrderRepository()
	os := orders.NewService(or)
	return newOrderController(os)
}

func TestPostOrder(t *testing.T) {
	oc := newTestOrdersHandler()

	req := test.NewRequest(t, "POST", "/order", nil)
	rr := test.ServeRequest(oc.PostOrder, req)

	t.Run("I should see a 201 response and a ID of 1", func(t *testing.T) {
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}

		expected := `{"order_id":1}`
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("postURL() returned an error: %s", rr.Body.String())
		}
	})

	return
}
