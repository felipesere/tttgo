package main

type Game struct {
	players [2]Player
	board   Board
}

func NewGame(players [2]Player, initialBoard Board) Game {
	return Game{
		players: players,
		board:   initialBoard,
	}
}

func (game *Game) RunToEnd() {
	for isNotOver(game) {
		currentPlayer := game.players[0]
		otherPlayer := game.players[1]
		game.board = currentPlayer.MakeMove(game.board)

		game.players[0] = otherPlayer
		game.players[1] = currentPlayer
	}
}

func isNotOver(game *Game) bool {
	return len(game.board.GetPossibleMoves()) > 0 && game.board.Winner() == nil
}
