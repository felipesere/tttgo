package main

import (
	"fmt"
)

type Line = []Cell

type Board struct {
	size  int
	cells []Cell
}

type Cell struct {
	mark  Mark
	move  int
	empty bool
}

func aMove(move int) Cell {
	return Cell{
		move:  move,
		empty: true,
	}
}

func takenBy(mark Mark) Cell {
	return Cell{
		mark:  mark,
		empty: false,
	}
}

func NewBoard(size int) Board {
	cells := make([]Cell, size*size)
	for idx := range cells {
		cells[idx] = aMove(idx)
	}

	return Board{
		size:  size,
		cells: cells,
	}
}

func (board *Board) GetPossibleMoves() []int {
	result := make([]int, 0)

	for _, cell := range board.cells {
		if cell.empty {
			result = append(result, cell.move)
		}
	}

	return result
}

func (board Board) IsFull() bool {
	return len(board.GetPossibleMoves()) == 0
}

func (board Board) MakeMove(move int, mark Mark) (Board, error) {
	if !board.IsMovePossible(move) {
		return board, fmt.Errorf("position %d was already taken", move)
	}

	other := Board{
		size:  board.size,
		cells: make([]Cell, len(board.cells)),
	}

	copy(other.cells, board.cells)

	other.cells[move] = takenBy(mark)

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

func (board *Board) Winner() *Mark {
	lines := make([]Line, 0)

	for idx := 0; idx < board.size; idx++ {
		lines = append(lines, rowAt(board, idx))
		lines = append(lines, columnAt(board, idx))
	}

	lines = append(lines, diagonals(board)...)

	return findWinner(lines)
}

func (board *Board) Rows() []Line {
	lines := make([]Line, board.size)

	for idx := 0; idx < board.size; idx++ {
		lines[idx] = rowAt(board, idx)
	}

	return lines
}

func rowAt(board *Board, number int) Line {
	beginning := number * board.size
	end := beginning + board.size
	return board.cells[beginning:end]
}

func columnAt(board *Board, number int) Line {
	column := make([]Cell, 0)

	for idx := 0; idx < board.size; idx++ {
		realIndex := number + idx*board.size
		column = append(column, board.cells[realIndex])
	}

	return column
}

func diagonals(board *Board) []Line {
	diagonal_one := make(Line, 0)
	diagonal_two := make(Line, 0)

	for idx := 0; idx < board.size; idx++ {
		x := idx + idx*board.size
		y := (idx + 1) * (board.size - 1)
		diagonal_one = append(diagonal_one, board.cells[x])
		diagonal_two = append(diagonal_two, board.cells[y])
	}

	return []Line{diagonal_one, diagonal_two}
}

func findWinner(lines []Line) *Mark {
	for _, line := range lines {
		winner := hasWinner(line)
		if winner != nil {
			return winner
		}
	}
	return nil
}

func hasWinner(line Line) *Mark {
	first := line[0]

	if first.empty {
		return nil
	}
	for _, cell := range line {
		if cell.mark != first.mark {
			return nil
		}
	}
	return &first.mark
}
