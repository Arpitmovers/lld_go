package main

import "time"

const (
	hourPrice = 10
	mintPrice = 3
)

type Ticket struct {
	startTime time.Time
	endTime   time.Time
	totalCost int
	vechile   Vechile
}

func NewTicket(v Vechile) *Ticket {
	return &Ticket{
		startTime: time.Now(),
		totalCost: 0,
		vechile:   v,
		endTime:   time.Time{},
	}
}

func (t *Ticket) SetExitTime() {
	t.endTime = time.Now()
}

func (t *Ticket) CalculateCost() float64 {

	duration := t.endTime.Sub(t.startTime)
	hours := duration.Hours()
	mints := duration.Minutes()

	remaingMints := mints / hours

	cost := hours*hourPrice + remaingMints*mintPrice

	return cost

}
