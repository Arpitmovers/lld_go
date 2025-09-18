package main

import (
	"fmt"
	"math/rand"
)

type Board struct {
	size        int
	board       [][]int
	ladderCount int
	snakeMap    map[int]int // mouth==>tail
	ladderMap   map[int]int // starting position ==> end position
	snakeCnt    int
}

func InitAll(boardSize int, ladderCnt int, snakeCnt int) *Board {
	b := &Board{
		size:        boardSize,
		ladderCount: ladderCnt,
		snakeCnt:    snakeCnt,
	}

	b.initBoard(boardSize)
	b.setLadders(ladderCnt)
	b.setSnake(snakeCnt)
	return b

}

// set  nos from 1 to n*n -1
func (b *Board) initBoard(n int) {
	no := 1
	b.board = make([][]int, n)

	for i := 0; i < n; i++ {
		b.board[i] = make([]int, n)
		for j := 0; j < n; j++ {
			b.board[i][j] = no
			no++
		}

	}
}

func (b *Board) PrintBoard() {

	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			fmt.Printf("%d ", b.board[i][j])
		}
		fmt.Println()
	}
}

// n is
func (b *Board) setLadders(ladderCnt int) {

	// randomly select [i][j] start , end and add to map
	//n := b.size * b.size
	n := 100
	i := 0
	for i < ladderCnt {

		p1 := rand.Intn(n + 1)
		p2 := rand.Intn(n + 1)

		start := p1
		end := p2

		if p1 > p2 {
			end = p2
			start = p1
		}

		if _, exists := b.ladderMap[start]; exists {
			continue
		} else {
			b.ladderMap[start] = end
			i++
		}

	}

}

func (b *Board) setSnake(snakeCnt int) {

	n := b.size * b.size
	i := 0
	for i < snakeCnt {

		p1 := rand.Intn(n + 1)
		p2 := rand.Intn(n + 1)

		start := p2
		end := p1

		if p1 > p2 {
			end = p2
			start = p1
		}

		if _, exists := b.snakeMap[start]; exists {
			continue
		} else {
			b.snakeMap[start] = end
			i++
		}

	}

}

func (b *Board) checkSnake(cell int) int {

	if dest, exists := b.snakeMap[cell]; exists {
		return dest
	}
	return -1
}

func (b *Board) checkLadder(cell int) int {

	if dest, exists := b.ladderMap[cell]; exists {
		return dest
	}
	return -1
}
