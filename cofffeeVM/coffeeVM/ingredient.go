package Cofffeevm

import "sync"

type Ingredient struct {
	Name     string
	Quantity int
	mu       sync.Mutex
	//Price int
}

func NewIngredient(name string, quantity int) *Ingredient {
	return &Ingredient{
		Name:     name,
		Quantity: quantity,
	}
}

func (ing *Ingredient) GetName() string {
	return ing.Name
}

func (ing *Ingredient) GetQuantity() int {
	ing.mu.Lock()
	defer ing.mu.Unlock()
	return ing.Quantity
}

func (ing *Ingredient) UpdateQuantity(quan int) {
	ing.mu.Lock()
	defer ing.mu.Unlock()
	ing.Quantity = 0
	// ing.Quantity = ing.Quantity + quan
}
