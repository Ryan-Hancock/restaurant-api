package orders

// Order repersents the ordering of multiple Items.
type Order struct {
	ID     int     `json:"id"`
	Lines  *[]Line `json:"lines"`
	IsPaid bool    `json:"is_paid"`
}

// Line repersents the indiviual purchase entries on an Order.
type Line struct {
	ID      int `json:"id"`
	ItemID  int `json:"item_id"`
	OrderID int `json:"order_id"`
}

// Bill repersents the costing for an Order
type Bill struct {
	ID      int `json:"id"`
	OrderID int `json:"order_id"`
}
