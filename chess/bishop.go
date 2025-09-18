package chess

import "math"

type Bishop struct {
	BasePiece
}

func NewBishop(isWhiteFlg bool) *Bishop {

	return &Bishop{
		BasePiece: BasePiece{
			Killed: false,
			White:  isWhiteFlg,
		},
	}
}

func (b *Bishop) CanMove(start Box, end Box, board *Board) bool {
	// Safety: prevent nil dereference
	if start.Piece == nil {
		return false
	}
	if end.Piece != nil && start.Piece.IsWhite() == end.Piece.IsWhite() {
		return false
	}

	// Valid bishop move only if |Δx| == |Δy|
	xDiff := math.Abs(float64(start.X - end.X))
	yDiff := math.Abs(float64(start.Y - end.Y))

	if xDiff != yDiff {
		return false
	}

	// You can also add path-blocking logic here if needed

	return true
}
