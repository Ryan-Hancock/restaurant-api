package memory

import (
	"github.com/ryan-hancock/resturant-api/pkg/items"
	"github.com/ryan-hancock/resturant-api/pkg/orders"
)

func (r *orderRepository) GetOrder(ID int) (o orders.Order, err error) {
	if o, ok := r.ordersDB[ID]; !ok {
		return o, items.ErrNotFound
	}

	return o, nil
}

func (r *orderRepository) InsertOrder(o orders.Order) (int, error) {
	var newID = len(r.ordersDB) + 1
	o.ID = newID
	r.ordersDB[newID] = o

	return newID, nil
}

func (r *orderRepository) UpdateOrder(o orders.Order) error {
	panic("not implemented") // TODO: Implement
}

func (r *orderRepository) InsertLine(l orders.Line) (int, error) {
	if _, ok := r.ordersDB[l.OrderID]; !ok {
		return 0, orders.ErrNotFound
	}

	if _, ok := r.itemsDB[l.ItemID]; !ok {
		return 0, items.ErrNotFound
	}

	var newID = len(r.linesDB) + 1
	l.ID = newID
	r.linesDB[newID] = l

	return newID, nil
}

func (r *orderRepository) GetLinesByOrderID(ID int) ([]orders.Line, error) {
	var lines []orders.Line
	for _, l := range r.linesDB {
		if l.OrderID == ID {
			lines = append(lines, l)
		}
	}

	return lines, nil
}
