package main

func main() {
	p1 := &Player{Name: "arpit"}
	p2 := &Player{Name: "test"}

	g1 := NewGame(p1, p2)
	g1.Start()

}
