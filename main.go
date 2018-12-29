package main

import (
	"os"
)

func main() {
	tui := Tui{writer: os.Stdout}
	human := NewHuman(X, tui)
	computer := NewComputer(O)
	var (
		emptyBoard Board = NewBoard(3)
		game       Game  = NewGame([2]Player{human, computer}, emptyBoard)
	)

	game.RunToEnd()
	tui.show(game.board)
}
