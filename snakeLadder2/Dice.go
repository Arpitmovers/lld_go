package main

import "math/rand"

type Dice struct {
	sides int
	count int
}

func (d *Dice) rollDice() int {

	start := d.count
	max := d.count * d.sides
	return rand.Intn(max-start+1) + start
}

func NewDice(sides, count int) *Dice {
	return &Dice{sides: sides, count: count}
}
