package main

type ParkSpot interface {
	IsOccupied(v *Vechile) bool
	ParkVechile(v *Vechile) bool
	UnPark(v *Vechile) bool
}

type ParkingSpot struct {
	vechile  *Vechile
	occupied bool
	id       int
	price    int
}

func NewParkingSpot(priceVar int, index int) *ParkingSpot {
	return &ParkingSpot{
		vechile:  &Vechile{},
		occupied: false,
		price:    priceVar,
		id:       index,
	}
}

func (spot *ParkingSpot) IsOccupied() bool {
	return spot.occupied
}

func (spot *ParkingSpot) ParkVechile(v *Vechile) bool {
	spot.vechile = v
	return true
}

func (spot *ParkingSpot) UnPark(v *Vechile) bool {
	spot.vechile = nil
	spot.occupied = false
	return true
}
