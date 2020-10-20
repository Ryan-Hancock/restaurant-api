package items

// Item repersents a orderable product.
type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
