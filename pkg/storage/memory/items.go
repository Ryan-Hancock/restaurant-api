package memory

import "github.com/ryan-hancock/resturant-api/pkg/items"

func (r *itemRepository) GetItem(ID int) (i items.Item, err error) {
	i, ok := r.itemsDB[ID]
	if !ok {
		return i, items.ErrNotFound
	}

	return
}

func (r *itemRepository) GetAllItems() []items.Item {
	itms := []items.Item{}
	for _, value := range r.itemsDB {
		itms = append(itms, value)
	}

	return itms
}

func (r *itemRepository) InsertItem(i items.Item) (int, error) {
	var newID = len(r.itemsDB) + 1
	i.ID = newID
	r.itemsDB[newID] = i

	return newID, nil
}

func (r *itemRepository) UpdateItem(i items.Item) error {
	found, err := r.GetItem(i.ID)
	if err != nil {
		return err
	}

	r.itemsDB[found.ID] = i
	return nil
}
