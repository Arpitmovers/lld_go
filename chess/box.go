package chess

type Box struct {
	X     int
	Y     int
	Piece Piece
}

func createBox(x int, y int, piece Piece) *Box {

	return &Box{
		X:     x,
		Y:     y,
		Piece: piece,
	}
}
