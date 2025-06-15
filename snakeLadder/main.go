package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	id       int
	name     string
	Position int
	Finished bool
}

type Game struct {
	Players []*Player
	Dice    *Dice
	Board   *Board
}

type Board struct {
	Size      int
	SnakeMap  map[int]int
	LadderMap map[int]int
}

func (b *Board) startGame() {

}

type Dice struct {
	sides int
}

func (d *Dice) rollDice() int {
	max := d.sides
	return rand.Intn(max) + 1
}

func newBoard(size int) *Board {
	snakesMap := make(map[int]int)
	snakesMap[16] = 7
	snakesMap[34] = 2
	snakesMap[44] = 11
	snakesMap[23] = 1
	snakesMap[32] = 4

	ladderMap := make(map[int]int)
	snakesMap[11] = 71
	snakesMap[19] = 64
	snakesMap[23] = 44
	snakesMap[31] = 61
	snakesMap[42] = 81

	return &Board{
		Size:      size,
		SnakeMap:  snakesMap,
		LadderMap: ladderMap,
	}
}

func (g *Game) Start() {
	finishedCount := 0

	for finishedCount < len(g.Players) {

		// all players
		for _, player := range g.Players {

			if player.Finished {
				continue
			}
			// get random no and move the player to new position
			move := g.Dice.rollDice()
			newPostion := player.Position + move
			fmt.Println("dice ", move)
			fmt.Println("move for player ", player.name, "moving from ", player.Position, "to ", newPostion)

			player.Position = newPostion

			if newPostion > g.Board.Size {
				fmt.Println("invalid move")
				player.Position -= move
				continue
			}

			if endPosition, ok := g.Board.SnakeMap[player.Position]; ok {
				fmt.Println("snake bite for player ", player.id, "at position", newPostion)
				player.Position = endPosition
			} else if ladderEndPos, ok := g.Board.LadderMap[newPostion]; ok {
				fmt.Println("ladder moving  player ", player.id, "to position", ladderEndPos)
				player.Position = ladderEndPos
			}

			// check for snake or ladder

			if newPostion == g.Board.Size {
				fmt.Println("player ", player.id, "won")
				player.Finished = true
				finishedCount++
			}
		}

		fmt.Println("------------------")
		time.Sleep(1 * time.Second)

	}
}

func main() {
	b1 := newBoard(10)
	p1 := &Player{name: "Arpit"}
	p2 := &Player{name: "Chikki"}

	dice := &Dice{sides: 7}

	g1 := &Game{
		Board:   b1,
		Players: []*Player{p1, p2},
		Dice:    dice,
	}

	g1.Start()

	fmt.Println("starting game")

}
