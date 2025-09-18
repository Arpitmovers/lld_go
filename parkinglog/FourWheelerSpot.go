package main

type FourWheelSpot struct {
	*ParkingSpot
}

// type ParkSpot interface {
// 	IsOccupied(v *Vechile) bool
// 	ParkVechile(v *Vechile) bool
// 	UnPark(v *Vechile) bool
// }

func (spot *FourWheelSpot) ParkVechile(v *Vechile) bool {

	if spot.IsOccupied() {
		return false
	}

	if v.size == TwoWheel {
		return false
	}

	spot.vechile = v
	return true
}
