package main

import (
	"fmt"
)

type Mark string

type Line = []Mark

const (
	X     Mark = "X"
	O     Mark = "O"
	EMPTY Mark = "-"
)

type Board struct {
	size  int
	marks []Mark
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
	if !board.IsMovePossible(move) {
		return board, fmt.Errorf("position %d was already taken", move)
	}

	other := Board{
		size:  board.size,
		marks: make([]Mark, len(board.marks)),
	}

	copy(other.marks, board.marks)

	other.marks[move] = mark

	return other, nil
}

func (board *Board) IsMovePossible(target int) bool {
	for _, mark := range board.GetPossibleMoves() {
		if mark == target {
			return true
		}
	}
	return false
}

func (board *Board) Winner() Mark {
	lines := make([][]Mark, 0)

	for idx := 0; idx < board.size; idx++ {
		lines = append(lines, rowAt(board, idx))
		lines = append(lines, columnAt(board, idx))
	}

	lines = append(lines, diagonals(board)...)

	return findWinner(lines)
}

func rowAt(board *Board, number int) Line {
	beginning := number * board.size
	end := beginning + board.size
	return board.marks[beginning:end]
}

func columnAt(board *Board, number int) Line {
	column := make([]Mark, 0)

	for idx := 0; idx < board.size; idx++ {
		realIndex := number + idx*board.size
		column = append(column, board.marks[realIndex])
	}

	return column
}

func diagonals(board *Board) []Line {
	diagonal_one := make([]Mark, 0)
	diagonal_two := make([]Mark, 0)

	for idx := 0; idx < board.size; idx++ {
		x := idx + idx*board.size
		y := (idx + 1) * (board.size - 1)
		diagonal_one = append(diagonal_one, board.marks[x])
		diagonal_two = append(diagonal_two, board.marks[y])
	}

	return [][]Mark{diagonal_one, diagonal_two}
}

func findWinner(lines []Line) Mark {
	for _, line := range lines {
		winner := hasWinner(line)
		if winner != nil {
			return *winner
		}
	}
	return EMPTY
}

func hasWinner(line Line) *Mark {
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
