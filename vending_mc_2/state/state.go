package state

// core abstraction. It defines the behavior expected from all concrete states.
type State interface {
	InsertMoney(int) error
	SelectItem(string) error
	DispenseItem() error
	Cancel() error
}
