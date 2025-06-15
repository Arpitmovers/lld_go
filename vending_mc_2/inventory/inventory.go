package inventory

import (
	"errors"
	"vendingMachine/model"
)

type Inventory struct {
	items map[string]*model.Item
}

// factroy to create inventroyt
func NewInvetory() *Inventory {
	return &Inventory{
		items: make(map[string]*model.Item),
	}
}

// add item
func (i *Inventory) AddItem(item model.Item) {
	i.items[item.Code] = &item
}

// getItem
func (i *Inventory) GetItem(code string) (*model.Item, error) {

	if item, ok := i.items[code]; ok && item.Quantity > 0 {
		return item, nil
	}
	return nil, errors.New("item not exist")

}

func (i *Inventory) DeductItem(code string) error {
	item, err := i.items[code]
	if err {
		return errors.New("uitem doenst exist")
	}
	if item.Quantity < 1 {
		return errors.New("item is over")
	}

	item.Quantity--
	return nil

}
