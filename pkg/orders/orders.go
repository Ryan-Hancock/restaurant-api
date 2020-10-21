package orders

// Order represents the ordering of multiple Items.
type Order struct {
	ID     int     `json:"id"`
	Lines  *[]Line `json:"lines"`
	IsPaid bool    `json:"is_paid"`
}

// Line represents the individual purchase entries on an Order.
type Line struct {
	ID      int `json:"id"`
	ItemID  int `json:"item_id"`
	OrderID int `json:"order_id"`
}