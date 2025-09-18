package main

type ParkingFloor struct {
	floorNo        int
	parkingSpots   map[VechileType]map[int]*ParkingSpot
	availableSpots int
}

// methods specific to theparkingfloor

// create ParkingFloor based on  floorNo ,noof spots

func NewParkingFloor(no int, twoWheelerCnt int, fourWheelerCnt int) *ParkingFloor {

	allSpots := make(map[VechileType]map[int]*ParkingSpot)

	if twoWheelerCnt > 0 {
		allSpots[TwoWheel] = CreateSpots(twoWheelerCnt)
	}

	if fourWheelerCnt > 0 {
		allSpots[FourWheel] = CreateSpots(fourWheelerCnt)
	}
	return &ParkingFloor{
		availableSpots: 10,
		parkingSpots:   allSpots,
		floorNo:        no,
	}
}

func CreateSpots(spotCnt int) map[int]*ParkingSpot {

	spots := make(map[int]*ParkingSpot)
	for i := 0; i < spotCnt; i++ {
		spot := NewParkingSpot(10, i+1)
		spots[spot.id] = spot
	}
	return spots
}

//search parkingSpot  within the floor

//show all spots onthe given floor
