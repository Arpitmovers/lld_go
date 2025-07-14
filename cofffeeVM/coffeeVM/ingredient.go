package cofffeevm

type Ingredient struct {
	Name     string
	Quantity int
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
	return ing.Quantity
}

func (ing *Ingredient) UpdateQuantity(quan int) {
	ing.Quantity = ing.Quantity + quan
}
