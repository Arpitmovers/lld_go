package main

type Player struct {
	id       int
	name     string
	Position int
}

func NewPlayer(id int, name string) *Player {
	return &Player{id: id, name: name, Position: 0}
}

func (p *Player) PlayerWon() bool {
	return p.Position == 100
}
