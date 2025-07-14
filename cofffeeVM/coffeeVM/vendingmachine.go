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
	mu          sync.Mutex
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

// total ingredients m/c has
func (vm *VendingMachine) initIngredients() {
	vm.Ingredients["Coffee"] = NewIngredient("Coffee", 50)
	vm.Ingredients["Water"] = NewIngredient("Water", 50)
	vm.Ingredients["Milk"] = NewIngredient("Milk", 10)
}

// menu consits of <name, price, map of ingredeients (this is a recepie)>
func (vm *VendingMachine) initMenu() {

	espessoRecepie := make(map[*Ingredient]int)

	// we refer the global ingredient here
	espessoRecepie[vm.Ingredients["Coffee"]] = 1
	espessoRecepie[vm.Ingredients["Water"]] = 2

	espesso := NewCoffee("espesso", 10, espessoRecepie)

	latteRecipie := make(map[*Ingredient]int)
	latteRecipie[vm.Ingredients["Coffee"]] = 1
	latteRecipie[vm.Ingredients["Water"]] = 1
	latteRecipie[vm.Ingredients["Milk"]] = 2

	latte := NewCoffee("latte", 15, latteRecipie)

	vm.menu = append(vm.menu, espesso)
	vm.menu = append(vm.menu, latte)
}

func (vm *VendingMachine) displayStock() {
	for _, coffee := range vm.menu {

		name := coffee.GetName()
		//	_ := coffee.GetPrice()

		recipe := coffee.GetRecipie()
		fmt.Println("_____________", name, "_______________")
		for ingreMap, _ := range recipe {
			ingrName := ingreMap.GetName()
			fmt.Println(ingrName, "quantity", vm.Ingredients[ingrName])

		}

	}
}

func (vm *VendingMachine) DisplayMenu() {

	for _, coffee := range vm.menu {
		fmt.Println(coffee.GetName(), coffee.GetPrice())
	}
}

func (vm *VendingMachine) DisplayInventory() {
	fmt.Println("------ Current Inventory ------")
	for name, ingredient := range vm.Ingredients {
		fmt.Printf("%s: %d units\n", name, ingredient.GetQuantity())
	}
}

func (vm *VendingMachine) SelectCoffee(c string) *Coffee {
	vm.mu.Lock()
	defer vm.mu.Unlock()
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

// update ingredentiensfor VM
func (vm *VendingMachine) updateIngredenets(c *Coffee) {
	for ing, requiredQuantity := range c.GetRecipie() {
		fmt.Println("debug ", ing.GetName(), requiredQuantity)
		ing.UpdateQuantity(-requiredQuantity)
		//name := ing.Name

		//vm.Ingredients[name].UpdateQuantity(-requiredQuantity)
	}

}
