package main

type GameEngine struct {
	board   *Board
	players []Player
	winner  Player
	dice    *Dice
}

func NewGame(boardSize int, ladders int, snakeCnt int) *GameEngine {

	return &GameEngine{
		board:  InitAll(boardSize, ladders, snakeCnt),
		dice:   NewDice(6, 4),
		winner: Player{},
	}
}

func (g *GameEngine) AddPlayers(name string) {
	g.players = append(g.players, *NewPlayer(len(g.players), name))
}

/*
checks for winner
changes position of
*/
func (g *GameEngine) Play() {

	for {

		// get currnet player
		// get current position and roll dice
		// move the current player
		// check if player won --> end

		for i := 0; i < len(g.players); i++ {
			currentPlayerposition := g.players[i].Position
			incr := g.dice.rollDice()
			newPosition := currentPlayerposition + incr

			// check for SNAKE

			newPosition = g.board.checkLadder(newPosition)
			if newPosition != -1 {
				// move  player ro new position
				g.players[i].Position = newPosition
			} else {
				newPosition = g.board.checkSnake(newPosition)
				g.players[i].Position = newPosition
			}

			if g.players[i].PlayerWon() {
				break
			}
			g.board.PrintBoard()

		}

	}

}
