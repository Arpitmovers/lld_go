package cofffeevm

type payment struct {
	Amount int
}

func NewPayment(amount int) *payment {

	return &payment{Amount: amount}
}
