package store

import "github.com/ryan-hancock/resturant-api/resturant"

var (
	fakeItemDB  map[string]*resturant.Item
	fakeOrderDB map[string]*resturant.Order
)

type ItemService struct{}

func (i *ItemService) List() ([]*resturant.Item, error) {
	panic("not implemented") // TODO: Implement
}

func (i *ItemService) Get(ID int) (*resturant.Item, error) {
	panic("not implemented") // TODO: Implement
}

func (i *ItemService) Update(_ *resturant.Item) error {
	panic("not implemented") // TODO: Implement
}

type OrderService struct{}

func (o *OrderService) Create() (*resturant.Order, error) {
	panic("not implemented") // TODO: Implement
}

func (o *OrderService) AddItem(orderID int, item *resturant.Item) error {
	panic("not implemented") // TODO: Implement
}

func (o *OrderService) Pay(orderID int) error {
	panic("not implemented") // TODO: Implement
}
