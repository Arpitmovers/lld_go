package chess



type Board struct {
	Boxes [][]*Box
}



func initiliseBoard() *Board {

	board =&Board{
		Boxes :make([][]Box , 8)
	}
	
	for i:= range board.Boxes {
		board.Boxes[i] = make([]Box,8)
	}
	placePieces(board);
	// add black Player
	return board
}

placePieces(board *Board){


	board.Boxes[0][0] = createBox(0,0,NewRook(true) )
	board.Boxes[0][1] = createBox(0,1,NewKnight(true) )
	board.Boxes[0][2] = createBox(0,2,NewBishop(true) )
	board.Boxes[0][3] = createBox(0,3,NewQueen(true) )
	board.Boxes[0][4] = createBox(0,4,NewKing(true) )
	board.Boxes[0][5] = createBox(0,5,NewBishop(true) )
	board.Boxes[0][6] = createBox(0,6,NewRook(true) )
	board.Boxes[0][7] = createBox(0,7,NewRook(true) )


	// add all pawsn for white
	for i:=0;i<8;i++{
		board.Boxes[1][i] = createBox(1,i,NewPawn(true))
	}




	board.Boxes[7][0] = createBox(7,0,NewRook(true) )
	board.Boxes[7][1] = createBox(7,1,NewKnight(true) )
	board.Boxes[7][2] = createBox(7,2,NewBishop(true) )
	board.Boxes[7][3] = createBox(7,3,NewQueen(true) )
	board.Boxes[7][4] = createBox(7,4,NewKing(true) )
	board.Boxes[7][5] = createBox(7,5,NewBishop(true) )
	board.Boxes[7][6] = createBox(7,6,NewRook(true) )
	board.Boxes[7][7] = createBox(7,7,NewRook(true) )


		// add all pawsn for white
		for i:=0;i<8;i++{
			board.Boxes[7][i] = createBox(,i,NewPawn(true))
		}
	

}
