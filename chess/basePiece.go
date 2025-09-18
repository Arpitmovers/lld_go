package chess

type BasePiece struct {
	Killed bool
	White  bool
}

func (p *BasePiece) SetKilled(killed bool) {
	p.Killed = killed
}

func (p *BasePiece) IsKilled() bool {
	return p.Killed
}

func (p *BasePiece) IsWhite() bool {
	return p.White
}
