package chess

type King struct {
	BasePiece
}

func NewKing(whiteFlag bool) *King {

	return &King{
		BasePiece: BasePiece{
			White:  whiteFlag,
			Killed: false,
		},
	}
}

// func (k *King) CanMove(start Box, end Box, b *Board) bool {
// 	xDiff := math.Abs(float64(start.X - end.X))

// 	yDiff := math.Abs(float64(start.Y - end.Y))

// 	if xDiff <= 1 && yDiff <= 1 {
// 		return true
// 	}
// 	return false
// }
