package items

import "errors"

// ErrNotFound for when an Item can not found.
var ErrNotFound = errors.New("item not found")

// ErrStorage for when the Item storage fails on operation.
var ErrStorage = errors.New("item storage failed")

// Repository provides access to the Item storage.
type Repository interface {
	GetItem(int) (Item, error)
	GetAllItems() []Item
	InsertItem(Item) (int, error)
	UpdateItem(Item) error
}

// Service provides Item operations.
type Service interface {
	NewItem(Item) (int, error)
	GetItem(ID int) (Item, error)
	GetItems() []Item
	ChangeItemPrice(ID int, price float32) error
}

type service struct {
	r Repository
}

// NewService creates a new Item service.
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) NewItem(i Item) (int, error) {
	return s.r.InsertItem(i)
}

func (s *service) GetItem(ID int) (Item, error) {
	item, err := s.r.GetItem(ID)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (s *service) GetItems() []Item {
	return s.r.GetAllItems()
}

func (s *service) ChangeItemPrice(ID int, price float32) error {
	item, err := s.r.GetItem(ID)
	if err != nil {
		return err
	}

	item.Price = price
	return s.r.UpdateItem(item)
}
