package Cofffeevm

import "fmt"

func Run() {

	vm := GetVendingMaching()
	vm.DisplayMenu()

	// select coffee
	selectedCoffee := vm.SelectCoffee("espesso")
	payment := NewPayment(100)

	// take payment
	if err := vm.DispenseItem(selectedCoffee, payment); err != nil {
		fmt.Println("error", err)
	}

}
