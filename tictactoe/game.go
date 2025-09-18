package main

import "fmt"

type Game struct {
	board       *Board
	p1          *Player
	p2          *Player
	currentTurn *Player
}

func NewGame(pl1 *Player, pl2 *Player) *Game {
	return &Game{
		board:       NewBoard(),
		p1:          pl1,
		p2:          pl2,
		currentTurn: pl1,
	}
}

func (g *Game) Start() {
	for !g.isWon() && g.IsFull() {

		// accpet user input for the moves
		row := 1
		col := 3
		g.board.MakeMove(row, col, '-')

		g.board.printBoard()
		g.changePlayer()
	}
	if g.isWon() {
		fmt.Println("player won is", g.currentTurn)
		return
	}
	fmt.Println("game is drawn")

}

func (g *Game) isWon() bool {

	// check all rows
	for i := 0; i < 3; i++ {
		currRow := g.board.grid[i]

		if currRow[0] != '-' && currRow[0] == currRow[1] && currRow[1] == currRow[2] {
			return true
		}
	}
	for j := 0; j < 3; j++ {
		if g.board.grid[0][j] != '-' && g.board.grid[0][j] == g.board.grid[1][j] && g.board.grid[1][j] == g.board.grid[2][j] {
			return true
		}
	}

	// diagonal
	if g.board.grid[0][0] != '-' && g.board.grid[0][0] == g.board.grid[1][1] && g.board.grid[1][1] == g.board.grid[2][2] {
		return true
	}

	if g.board.grid[0][2] != '-' && g.board.grid[0][2] == g.board.grid[1][1] && g.board.grid[1][1] == g.board.grid[0][3] {
		return true
	}

	return false
}

func (g *Game) IsFull() bool {
	return false
}

func (g *Game) changePlayer() {
	if g.currentTurn == g.p1 {
		g.currentTurn = g.p2
		return
	}
	g.currentTurn = g.p1
}
