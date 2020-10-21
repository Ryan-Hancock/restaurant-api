package items

import (
	"testing"
)

func getTestItemService() Service {
	ms := &mockStorage{
		store: make(map[int]Item),
	}

	return &service{ms}
}

func Test_service_NewItem(t *testing.T) {
	s := getTestItemService()
	i1 := Item{
		ID:    1,
		Name:  "burger",
		Price: 1.99,
	}

	if _, err := s.NewItem(i1); err != nil {
		t.Errorf("got an error while inserting, %v", err)
	}

	if length := len(s.GetItems()); length != 1 {
		t.Errorf("wanted to only see 1 but got %d", length)
	}
}

func Test_service_GetItem(t *testing.T) {
	s := getTestItemService()
	i1 := Item{
		ID:    1,
		Name:  "burger",
		Price: 1.99,
	}
	s.NewItem(i1)

	item, err := s.GetItem(i1.ID)
	if err != nil {
		t.Errorf("got an error while getting, %v", err)
	}

	if item.ID != i1.ID {
		t.Errorf("wrong item came back from store, %v", item)
	}
}

func Test_service_ChangeItemPrice(t *testing.T) {
	s := getTestItemService()
	i1 := Item{
		ID:    1,
		Name:  "burger",
		Price: 1.99,
	}
	s.NewItem(i1)

	err := s.ChangeItemPrice(i1.ID, 2.99)
	if err != nil {
		t.Errorf("got an error while updating, %v", err)
	}

	item, _ := s.GetItem(i1.ID)
	if item.Price != 2.99 {
		t.Errorf("expected 2.99 but got, %v", item.Price)
	}
}

type mockStorage struct {
	store map[int]Item
}

func (m *mockStorage) GetItem(i int) (Item, error) {
	item, _ := m.store[i]
	return item, nil
}

func (m *mockStorage) GetAllItems() []Item {
	items := []Item{}

	for _, itm := range m.store {
		items = append(items, itm)
	}

	return items
}

func (m *mockStorage) InsertItem(item Item) (int, error) {
	m.store[item.ID] = item
	return item.ID, nil
}

func (m *mockStorage) UpdateItem(item Item) error {
	m.store[item.ID] = item
	return nil
}
