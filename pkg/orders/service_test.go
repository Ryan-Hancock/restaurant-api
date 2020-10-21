package orders

import (
	"testing"

	"github.com/ryan-hancock/restaurant-api/pkg/items"
)

func getTestOrderService() Service {
	is := make(map[int]items.Item)
	is[1] = items.Item{ID: 1, Name: "Burger", Price: 5.99}

	ms := &mockStorage{
		storeOrder: make(map[int]Order),
		storeLine:  make(map[int]Line),
		storeItem:  is,
	}

	return &service{ms}
}

func Test_service_NewOrder(t *testing.T) {
	s := getTestOrderService()
	_, err := s.NewOrder()
	if err != nil {
		t.Errorf("got an error while creating, %v", err)
	}

	id1, _ := s.NewOrder()
	if id1 == 0 {
		t.Errorf("no id was set, got %v", id1)
	}
}

func Test_service_Pay(t *testing.T) {
	s := getTestOrderService()
	id, _ := s.NewOrder()

	_, err := s.AppendLine(Line{OrderID: id, ItemID: 1})
	if err != nil {
		t.Errorf("got an error while appending, %v", err)
	}

	err = s.Pay(id, 4.99)
	if err == nil {
		t.Errorf("we should have failed here, %v", err)
	}

	err = s.Pay(id, 5.99)
	if err != nil {
		t.Errorf("got an error while paying, %v", err)
	}

	order, _ := s.GetOrderWithLines(id)
	if err != nil {
		t.Errorf("got an error while getting order with lines, %v", err)
	}

	if !order.IsPaid {
		t.Error("expected to see true for isPaid")
	}
}

func Test_service_GetOrderWithLines(t *testing.T) {
	s := getTestOrderService()
	id, _ := s.NewOrder()

	_, err := s.AppendLine(Line{OrderID: id, ItemID: 1})
	if err != nil {
		t.Errorf("got an error while appending, %v", err)
	}

	order, err := s.GetOrderWithLines(id)
	if err != nil {
		t.Errorf("got an error while getting order with lines, %v", err)
	}

	if order.ID != id {
		t.Errorf("expected the same id, %v but got %v", id, order.ID)
	}

	if len(*order.Lines) != 1 {
		t.Errorf("expected the right length of line, %v but got %v", 1, len(*order.Lines))
	}
}

type mockStorage struct {
	storeOrder map[int]Order
	storeLine  map[int]Line
	storeItem  map[int]items.Item
}

func (m *mockStorage) GetOrder(i int) (Order, error) {
	return m.storeOrder[i], nil
}

func (m *mockStorage) InsertOrder(o Order) (int, error) {
	oID := len(m.storeOrder) + 1
	o.ID = oID
	m.storeOrder[oID] = o
	return oID, nil
}

func (m *mockStorage) UpdateOrder(o Order) error {
	m.storeOrder[o.ID] = o
	return nil
}

func (m *mockStorage) InsertLine(l Line) (int, error) {
	m.storeLine[l.ID] = l
	return l.ID, nil
}

func (m *mockStorage) GetLinesByOrderID(i int) ([]Line, error) {
	var lines []Line
	for _, l := range m.storeLine {
		if l.OrderID == i {
			lines = append(lines, l)
		}
	}

	return lines, nil
}

func (m *mockStorage) GetLinesPrice(ID int) (float32, error) {
	lines, _ := m.GetLinesByOrderID(ID)
	var amount float32
	for _, l := range lines {
		amount = amount + m.storeItem[l.ItemID].Price
	}

	return amount, nil
}
