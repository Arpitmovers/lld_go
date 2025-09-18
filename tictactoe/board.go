package main

import "fmt"

type Board struct {
	grid [3][3]rune
}

func NewBoard() *Board {
	b := &Board{}
	b.initGrid()
	return b
}

func (b *Board) initGrid() {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			b.grid[row][col] = '_'

		}
	}
}

func (b *Board) MakeMove(row int, col int, symbol rune) error {

	b.grid[row][col] = symbol

}

func (b *Board) printBoard() {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			//b.grid[row][col] = '_'
			fmt.Printf("%c ", b.grid[row][col])

		}
		fmt.Println()
	}
	fmt.Println()
}
