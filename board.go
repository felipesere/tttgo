package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type Mark string

const (
	X     Mark = "X"
	O     Mark = "O"
	EMPTY Mark = "-"
)

type Board struct {
	size  int
	marks []Mark
}

func (board *Board) isMovePossible(target int) bool {
	for _, mark := range board.GetPossibleMoves() {
		if mark == target {
			return true
		}
	}
	return false
}

func NewBoard(size int) Board {
	marks := make([]Mark, size*size)
	for idx := range marks {
		marks[idx] = EMPTY
	}

	return Board{
		size:  size,
		marks: marks,
	}
}

func (board *Board) GetPossibleMoves() []int {
	result := make([]int, 0)

	for idx, mark := range board.marks {
		if mark == EMPTY {
			result = append(result, idx)
		}
	}

	return result
}

func (board Board) MakeMove(move int, mark Mark) (Board, error) {
	if !board.isMovePossible(move) {
		return board, fmt.Errorf("position %d was already taken", move)
	}

	other := Board{}
	err := copier.Copy(&other, &board)

	other.marks[move] = mark

	return other, err
}

func (board *Board) Winner() Mark {
	lines := make([][]Mark, 0)

	// [ 0 1 2
	//   3 4 5
	//   6 7 8 ]

	for idx := 0; idx < board.size; idx++ {
		lines = append(lines, rowAt(board, idx))
		lines = append(lines, columnAt(board, idx))
	}

	return findWinner(lines)
}

func rowAt(board *Board, number int) []Mark {
	beginning := number * board.size
	end := beginning + board.size
	return board.marks[beginning:end]
}

func columnAt(board *Board, number int) []Mark {
	column := make([]Mark, 0)

	for idx := 0; idx < board.size; idx++ {
		realIndex := number + idx*board.size
		column = append(column, board.marks[realIndex])
	}

	return column
}

func findWinner(lines [][]Mark) Mark {
	for _, line := range lines {
		winner := hasWinner(line)
		if winner != nil {
			return *winner
		}
	}
	return EMPTY
}

func hasWinner(line []Mark) *Mark {
	first := line[0]

	if first == EMPTY {
		return nil
	}
	for _, mark := range line {
		if mark != first {
			return nil
		}
	}
	return &first
}
