package chess

// elephant
type Rook struct {
	BasePiece
}

func NewRook(whiteFlag bool) *Rook {
	return &Rook{
		BasePiece: BasePiece{
			White:  true,
			Killed: false,
		},
	}
}

func (r *Rook) CanMove(start Box, end Box) bool {

	if start.X != end.X && start.Y != end.Y {
		return false
	}

	// Additional: Check for same color piece blocking
	if end.Piece != nil && end.Piece.IsWhite() == r.IsWhite() {
		return false
	}

	return true
}
