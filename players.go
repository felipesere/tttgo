package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Player interface {
	MakeMove(board Board) Board
}

type Computer struct {
	mark Mark
}

type Human struct {
	mark    Mark
	display Display
}

func (human Human) MakeMove(board Board) Board {
	moves := board.GetPossibleMoves()

	move := moves[rand.Intn(len(moves))]

	result, _ := board.MakeMove(move, human.mark)

	return result
}

func NewHuman(mark Mark, display Display) Player {
	return Human{
		mark:    mark,
		display: display,
	}
}

func NewComputer(mark Mark) Player {
	return Computer{
		mark: mark,
	}
}

func (computer *Computer) Opponent() Computer {
	var other Mark
	if computer.mark == X {
		other = O
	} else {
		other = X
	}

	return Computer{
		mark: other,
	}
}

type ScoredMove struct {
	move  int
	score int
}

func (computer Computer) MakeMove(board Board) Board {
	move := computer.bestMove(board)
	newBoard, _ := board.MakeMove(move, computer.mark)

	return newBoard
}

func (computer Computer) bestMove(board Board) int {
	possibleMoves := board.GetPossibleMoves()
	bestScore := ScoredMove{score: -1000, move: possibleMoves[0]}

	for _, move := range possibleMoves {
		possibleBoard, _ := board.MakeMove(move, computer.mark)

		score := scoreBoard(possibleBoard, computer)

		if score > bestScore.score {
			bestScore = ScoredMove{move: move, score: score}
		}
	}

	return bestScore.move
}

func scoreBoard(board Board, self Computer) int {
	amountOfRemainingMoves := len(board.GetPossibleMoves())
	opponent := self.Opponent()
	possibleWinner := board.Winner()

	if self.mark.Is(possibleWinner) {
		return amountOfRemainingMoves
	} else if opponent.mark.Is(possibleWinner) {
		return -10 + amountOfRemainingMoves
	} else if board.IsFull() {
		return 0
	}

	bestOpponentBoard := opponent.MakeMove(board)

	return -1 * scoreBoard(bestOpponentBoard, opponent)
}
