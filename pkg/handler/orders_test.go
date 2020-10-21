package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/ryan-hancock/resturant-api/pkg/items"
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
	rr := test.ServeRequest("/order", oc.PostOrder, req)

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

func TestPatchOrderWithItem(t *testing.T) {
	oc := newTestOrdersHandler()
	ic := newTestItemHandler()

	oid, _ := oc.s.NewOrder()
	iid, _ := ic.s.NewItem(items.Item{Name: "burger"})

	t.Run("I should see a 204 response with orderID of 1 and itemID of 1", func(t *testing.T) {
		req := test.NewRequest(t, "PATCH", fmt.Sprintf("/order/%d/additem/%d", oid, iid), nil)
		rr := test.ServeRequest("/order/{orderID}/additem/{itemID}", oc.PatchOrderWithItem, req)

		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNoContent)
		}
	})

	t.Run("I should see a 404 response with orderID of 1 and itemID of 3 as it does not exisit", func(t *testing.T) {
		req := test.NewRequest(t, "PATCH", fmt.Sprintf("/order/%d/additem/%d", oid, 3), nil)
		rr := test.ServeRequest("/order/{orderID}/additem/{itemID}", oc.PatchOrderWithItem, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNotFound)
		}
	})

	t.Run("I should be able to add multiple of the same item 1", func(t *testing.T) {
		req := test.NewRequest(t, "PATCH", fmt.Sprintf("/order/%d/additem/%d", oid, iid), nil)
		rr := test.ServeRequest("/order/{orderID}/additem/{itemID}", oc.PatchOrderWithItem, req)

		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNoContent)
		}

		orders, _ := oc.s.GetOrderWithLines(oid)
		var count int
		for _, l := range *orders.Lines {
			if l.ItemID == iid {
				count++
			}
		}

		if count != 2 {
			t.Errorf("get lines for order should of: got %v want %v",
				count, 2)
		}
	})
}

func TestPostOrderPay(t *testing.T) {
	oc := newTestOrdersHandler()
	ic := newTestItemHandler()

	oid, _ := oc.s.NewOrder()
	iid, _ := ic.s.NewItem(items.Item{Name: "burger", Price: 5.99})
	oc.s.AppendLine(orders.Line{ItemID: iid, OrderID: oid})

	t.Run("I should see a 400 response if I under pay", func(t *testing.T) {
		var jsonStr = []byte(`{"amount":4.99}`)
		req := test.NewRequest(t, "POST", fmt.Sprintf("/order/%d/pay", oid), bytes.NewBuffer(jsonStr))
		rr := test.ServeRequest("/order/{orderID}/pay", oc.PostOrderPay, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})

	t.Run("I should see a 204 response with orderID of 1 and itemID of 1", func(t *testing.T) {
		var jsonStr = []byte(`{"amount":5.99}`)
		req := test.NewRequest(t, "POST", fmt.Sprintf("/order/%d/pay", oid), bytes.NewBuffer(jsonStr))
		rr := test.ServeRequest("/order/{orderID}/pay", oc.PostOrderPay, req)

		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNoContent)
		}
	})
}
