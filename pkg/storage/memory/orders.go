package memory

import "github.com/ryan-hancock/resturant-api/pkg/orders"

func (r *orderRepository) GetOrder(ID int) (orders.Order, error) {
	panic("not implemented") // TODO: Implement
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
	panic("not implemented") // TODO: Implement
}

func (r *orderRepository) GetLinesByOrderID(ID int) ([]orders.Line, error) {
	panic("not implemented") // TODO: Implement
}
