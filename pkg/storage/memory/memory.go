package memory

import (
	"github.com/ryan-hancock/resturant-api/pkg/items"
	"github.com/ryan-hancock/resturant-api/pkg/orders"
)

var (
	itemsDB  = make(map[int]items.Item)
	ordersDB = make(map[int]orders.Order)
	linesDB  = make(map[int]orders.Line)
)

type itemRepository struct {
	itemsDB map[int]items.Item
	itemID  int
}

// NewItemRepository creates a new items.Repository
func NewItemRepository() items.Repository {
	return &itemRepository{
		itemsDB: itemsDB,
	}
}

type orderRepository struct {
	ordersDB map[int]orders.Order
	linesDB  map[int]orders.Line
	itemsDB  map[int]items.Item

	orderID int
}

// NewOrderRepository creates a new items.Repository
func NewOrderRepository() orders.Repository {
	return &orderRepository{
		ordersDB: ordersDB,
		itemsDB:  itemsDB,
		linesDB:  linesDB,
	}
}
