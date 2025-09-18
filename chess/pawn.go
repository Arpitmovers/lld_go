package chess

import "math"

type Pawn struct {
	BasePiece
}

func NewPawn(isWhite bool) *Pawn {
	return &Pawn{
		BasePiece: BasePiece{
			White:  isWhite,
			Killed: false,
		},
	}
}
func CanMove(start Box, end Box, b *Board) bool {

	if end.Piece.IsWhite() == start.Piece.IsWhite() {
		return false
	}

	// is diagonall and box is not occupied then we can place
	xDiff := math.Abs(float64(start.X - end.X))
	yDiff := math.Abs(float64(start.Y - end.Y))

	// no piece is placed
	if xDiff+yDiff == 1 && end.Piece. == nil {
		return true
	}

	// Kill and move diagonally  , same as queen
	if xDiff+yDiff == 1 && end.Piece.IsWhite() != start.Piece.IsWhite() {
		return true
	}
}


1256: 