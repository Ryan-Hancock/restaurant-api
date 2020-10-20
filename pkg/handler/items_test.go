package handler

import (
	"net/http"
	"testing"

	"github.com/ryan-hancock/resturant-api/pkg/items"
	"github.com/ryan-hancock/resturant-api/pkg/test"
	"github.com/stretchr/testify/mock"
)

type MockedItemRepo struct {
	mock.Mock
}

func (m *MockedItemRepo) InsertItem(items.Item) (int, error) {
	args := m.Called(items.Item{})
	return args.Int(0), args.Error(1)
}

func getItemController() *itemController {
	var r MockedItemRepo
	is := items.NewService(r)
	return newItemController(is)
}

func TestGetItems(t *testing.T) {
	ic := getItemController()
	req := test.NewRequest(t, "GET", "/item", nil)
	rr := test.ServeRequest(ic.GetItems, req)

	t.Run("I should see a 200 response", func(t *testing.T) {
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}
	})

	return
}

func TestPostItem(t *testing.T) {
	ic := getItemController()
	req := test.NewRequest(t, "POST", "/item", nil)
	rr := test.ServeRequest(ic.PostItem, req)

	t.Run("I should see a 200 response", func(t *testing.T) {
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}
	})

	return
}

func TestPatchItem(t *testing.T) {
	return
}
