package machine

import (
	"errors"
	"fmt"
)

// each of below struct implements State interfacee and encapsulates logic  specific to that state
type IdleState struct {
	VM *VendingMachine
}
type HasMoneyState struct {
	VM *VendingMachine
}

type DispenseItemState struct {
	VM *VendingMachine
}

func (s *IdleState) InsertMoney(amount int) error {
	s.VM.Balance += amount
	fmt.Println("Money inserted:", amount)
	s.VM.SetState(&HasMoneyState{VM: s.VM})
	return nil
}

func (state *IdleState) DispenseItem() error {
	return errors.New("insert money and select item")
}

func (state *IdleState) SelectItem(string) error {
	return errors.New("cant select item in idel state")
}

func (state *IdleState) Cancel() error {
	return errors.New("nothing to cancel")
}

////////////////////////////////////////////////////

func (s *HasMoneyState) InsertMoney(amount int) error {
	s.VM.Balance += amount
	fmt.Println("Additional money inserted:", amount)
	return nil
}

func (s *HasMoneyState) SelectItem(code string) error {
	item, err := s.VM.Inventory.GetItem(code)
	if err != nil {
		return err
	}
	if item.Price > s.VM.Balance {
		return errors.New("insufficient balance")
	}
	s.VM.SelectedItem = item
	s.VM.SetState(&DispenseItemState{VM: s.VM})
	return nil
}

func (s *HasMoneyState) DispenseItem() error {
	return errors.New("select item first")
}

func (s *HasMoneyState) Cancel() error {
	fmt.Printf("Transaction cancelled. Returning money: %d\n", s.VM.Balance)
	s.VM.Balance = 0
	s.VM.SetState(&IdleState{VM: s.VM})
	return nil
}

//////////////////////////////////////////////////////////////////

func (s *DispenseItemState) InsertMoney(m int) error {
	return errors.New("in despense state, cant insert timte")
}

func (s *DispenseItemState) SelectItem(code string) error {
	return errors.New("in despense state, cant select timte")
}

func (s *DispenseItemState) DispenseItem() error {
	item := s.VM.SelectedItem

	err := s.VM.Inventory.DeductItem(item.Code)
	if err != nil {
		return err
	}

	s.VM.Balance -= item.Price
	s.VM.SelectedItem = nil

	s.VM.SetState(&IdleState{VM: s.VM})
	return nil

}

func (s *DispenseItemState) Cancel() error {
	return errors.New("cannot cancel while dispensing")
}

//////////////////////////////////////////
