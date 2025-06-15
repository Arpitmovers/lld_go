package main

import (
	"vendingMachine/machine"
	"vendingMachine/model"
)

func main() {

	machine := machine.NewVendingMachine()

	machine.Inventory.AddItem(model.Item{Code: "apple", Name: "Red Apples", Price: 50, Quantity: 12})
	machine.Inventory.AddItem(model.Item{Code: "buttermilk", Name: "Butter Milk", Price: 12, Quantity: 100})

	_ = machine.InsertMoney(12)
	machine.SelectItem("apple")
	machine.DispenseItem()

}
