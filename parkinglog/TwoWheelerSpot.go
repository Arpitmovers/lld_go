package main

import "fmt"

type TwoWheelerSpot struct {
	ParkingSpot
}

// type ParkSpot interface {
// 	IsOccupied(v *Vechile) bool
// 	ParkVechile(v *Vechile) bool
// 	UnPark(v *Vechile) bool
// }

func (spot *TwoWheelerSpot) ParkVechile(v *Vechile) bool {

	if spot.occupied {
		return false
	}
	if v.size != TwoWheel {
		fmt.Println("wrong parking type")
		return false
	}
	spot.vechile = v
	return true
}

func (spot *TwoWheelerSpot) UnPark(v *Vechile) bool {

	if spot.occupied {
		spot.vechile = nil
		return true
	}

	return false
}
