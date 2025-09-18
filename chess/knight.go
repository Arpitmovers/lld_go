package chess

type Knight struct {
	BasePiece
}

func NewKnight(whiteFlag bool) *Knight {

	return &Knight{
		BasePiece: BasePiece{
			White:  whiteFlag,
			Killed: false,
		},
	}

}
