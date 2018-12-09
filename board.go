package main


import (
	"fmt"
	"github.com/jinzhu/copier"
)
type Mark string

const (
	X Mark = "X"
	O Mark = "O"
	EMPTY Mark = ""
)

type Board struct {
  marks [9]Mark
}

func (board *Board) isMovePossible(target int) bool {
    for _, mark := range board.GetPossibleMoves() {
    	if mark == target {
    		return true
		}
	}
    return false
}

func NewBoard() Board {
	return Board {
		marks: [9]Mark{EMPTY},
	}
}

func (board *Board) GetPossibleMoves() []int  {
	result := make([]int, 0)

	counter := 0
	for idx, mark := range board.marks {
		if mark == EMPTY {
			result = append(result, idx)
			counter += 1
		}
	}
	return result
}

func (board Board) MakeMove(move int, mark Mark) (Board, error) {
	if !board.isMovePossible(move) {
		return board, fmt.Errorf("position %d was already taken", move)
	}

    other := Board{}
	copier.Copy(&other, &board)

    other.marks[move] = mark

    return other, nil
}

