package machine

import (
	"vendingMachine/inventory"
	"vendingMachine/model"
	"vendingMachine/state"
)

type VendingMachine struct {
	Inventory    *inventory.Inventory
	CurrentState state.State // current Stagte chan be InsertMoney, SelectItem,DispenseItem, Cancel
	SelectedItem *model.Item
	Balance      int
}

func NewVendingMachine() *VendingMachine {
	vm := &VendingMachine{
		Inventory: inventory.NewInvetory(),
		Balance:   0,
	}
	idle := &IdleState{VM: vm}
	vm.SetState(idle)
	return vm
}

func (vm *VendingMachine) SetState(s state.State) {
	vm.CurrentState = s
}

func (vm *VendingMachine) InsertMoney(amount int) error {
	return vm.CurrentState.InsertMoney(amount)
}

func (vm *VendingMachine) SelectItem(code string) error {
	return vm.CurrentState.SelectItem(code)
}

func (vm *VendingMachine) DispenseItem() error {
	return vm.CurrentState.DispenseItem()
}

func (vm *VendingMachine) Cancel() error {
	return vm.CurrentState.Cancel()
}
