package Cofffeevm

type Coffee struct {
	Name    string
	Price   int
	Recipie map[*Ingredient]int
}

func NewCoffee(name string, price int, recipie map[*Ingredient]int) *Coffee {
	return &Coffee{
		Name:    name,
		Price:   price,
		Recipie: recipie,
	}
}

func (c *Coffee) GetName() string {
	return c.Name
}

func (c *Coffee) GetPrice() int {
	return c.Price
}

func (c *Coffee) GetRecipie() map[*Ingredient]int {
	return c.Recipie
}
