package chess

type Queen struct {
	BasePiece
}

func NewQueen(whiteFlag bool) *Queen {

	return &Queen{
		BasePiece: BasePiece{
			White:  whiteFlag,
			Killed: false,
		},
	}

}
