package resturant

// Order repersents the ordering of multiple Items.
type Order struct {
	ID     int     `json:"id"`
	Lines  []*Line `json:"lines"`
	IsPaid bool    `json:"is_paid"`
}

// Line repersents the indiviual purchase entries on an Order.
type Line struct {
	ID       int   `json:"id"`
	Item     *Item `json:"item_id"`
	Quantity int   `json:"quantity"`
}

// OrderService interface for interacting with Order.
type OrderService interface {
	Create() (*Order, error)
	AddItem(orderID int, item *Item) error
	Pay(orderID int) error
}

// Bill repersents the costing for an Order
type Bill struct {
	ID      int `json:"id"`
	OrderID int `json:"order_id"`
}
