package main

type player interface {
	MakeMove(board Board) Board
}

type Computer struct {
	mark Mark
}

func NewComputer(mark Mark) Computer {
	return Computer{
		mark: mark,
	}
}

func (computer *Computer) Opponent() Computer {
	var other Mark
	if computer.mark == X {
		other = O
	} else if computer.mark == O {
		other = X
	} else {
		other = EMPTY
	}

	return Computer{
		mark: other,
	}
}

type ScoredMove struct {
	move  int
	score int
}

func (computer *Computer) MakeMove(board Board) Board {
	move := computer.bestMove(board)
	newBoard, _ := board.MakeMove(move, computer.mark)

	return newBoard
}

func (computer *Computer) bestMove(board Board) int {
	scores := make([]ScoredMove, 0)
	for _, move := range board.GetPossibleMoves() {
		possibleBoard, _ := board.MakeMove(move, computer.mark)

		score := computer.scoreBoard(possibleBoard)

		scores = append(scores, ScoredMove{move: move, score: score})
	}

	return scores[0].move
}

func (computer *Computer) scoreBoard(board Board) int {
	possibleWinner := board.Winner()

	if possibleWinner == computer.mark {
		return 10
	} else if possibleWinner == computer.Opponent().mark {
		return -10
	}
	return 0
}
