package main

func main() {

	n := 10
	ladders := 4
	snakes := 4

	g1 := NewGame(n, ladders, snakes)
	g1.AddPlayers("Arpit")
	g1.AddPlayers("Saloni")
	g1.Play()
}
