package Cofffeevm

import "fmt"

func Run() {

	vm := GetVendingMaching()
	//	vm.DisplayMenu()

	// select coffee
	selectedCoffee := vm.SelectCoffee("espesso")
	payment := NewPayment(39)

	vm.DisplayInventory()

	// take payment
	if err := vm.DispenseItem(selectedCoffee, payment); err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("......................after dispensing.................................")
	vm.DisplayInventory()
}
