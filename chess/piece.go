package chess

type Piece interface {
	IsWhite() bool
	IsKilled() bool
	SetKilled(killed bool)
	CanMove(start, end *Box) bool
}
