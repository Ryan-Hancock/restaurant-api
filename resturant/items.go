package resturant

// Item repersents a orderable product.
type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

// ItemService interface for interacting with Item.
type ItemService interface {
	List() ([]*Item, error)
	Get(ID int) (*Item, error)
	Update(*Item) error
}
