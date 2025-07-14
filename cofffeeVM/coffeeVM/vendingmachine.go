package Cofffeevm

import (
	"fmt"
	"sync"
)

var (
	instance *VendingMachine
	once     sync.Once
)

type VendingMachine struct {
	menu        []*Coffee
	Ingredients map[string]*Ingredient
}

func GetVendingMaching() *VendingMachine {

	once.Do(func() {

		instance = &VendingMachine{
			menu:        make([]*Coffee, 0),
			Ingredients: make(map[string]*Ingredient),
		}
		instance.initIngredients()
		instance.initMenu()
	})
	return instance
}

func (vm *VendingMachine) initIngredients() {
	vm.Ingredients["Coffee"] = NewIngredient("coffee", 12)
	vm.Ingredients["Water"] = NewIngredient("Water", 12)
	vm.Ingredients["Milk"] = NewIngredient("Milk", 12)
}

func (vm *VendingMachine) initMenu() {
	espessoRecepie := make(map[*Ingredient]int)
	espessoRecepie[vm.Ingredients["Coffee"]] = 1
	espessoRecepie[vm.Ingredients["Water"]] = 2

	espesso := NewCoffee("espesso", 12, espessoRecepie)

	latteRecipie := make(map[*Ingredient]int)
	latteRecipie[vm.Ingredients["Coffee"]] = 1
	latteRecipie[vm.Ingredients["Water"]] = 1
	latteRecipie[vm.Ingredients["Milk"]] = 2

	vm.menu = append(vm.menu, espesso)
}

func (vm *VendingMachine) DisplayMenu() {

	for idx, coffee := range vm.menu {
		fmt.Println("idx ", idx, "val", coffee.GetName(), coffee.GetPrice())
	}
}

func (vm *VendingMachine) SelectCoffee(c string) *Coffee {

	for _, coffeeItem := range vm.menu {
		if c == coffeeItem.GetName() {
			return coffeeItem
		}
	}

	return nil
}

func (vm *VendingMachine) DispenseItem(c *Coffee, p *payment) error {

	if c == nil {
		return fmt.Errorf("invalid coffee selection")
	}

	if p.Amount < c.Price {
		return fmt.Errorf("insufficnet payment for %s", c.GetName())
	}

	result := vm.validateStockForPrep(c)
	if !result {
		fmt.Println("cant prepare ", c.Name)

	}
	// update ingredient stock that we have
	vm.updateIngredenets(c)

	change := p.Amount - c.Price

	if change > 0 {
		fmt.Println("take your change", change)
	}

	return nil
}

func (vm *VendingMachine) validateStockForPrep(c *Coffee) bool {

	for ing, requiredQuantity := range c.GetRecipie() {

		if ing.GetQuantity() < requiredQuantity {
			return false
		}
	}
	return true
}

func (vm *VendingMachine) updateIngredenets(c *Coffee) {
	for ing, requiredQuantity := range c.GetRecipie() {
		ing.UpdateQuantity(-requiredQuantity)
	}

}
