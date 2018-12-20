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

	if possibleWinner == self.mark {
		return amountOfRemainingMoves
	} else if possibleWinner == opponent.mark {
		return -10 + amountOfRemainingMoves
	} else if board.IsFull() {
		return 0
	}

	bestOpponentBoard := opponent.MakeMove(board)

	return -1 * scoreBoard(bestOpponentBoard, opponent)
}
