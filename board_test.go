package main

import (
	"testing"
)

func TestBoard_GetPossibleMoves(t *testing.T) {
	anEmptyBoard := NewBoard()

	numberOfPossibleMoves := len(anEmptyBoard.GetPossibleMoves())

	if numberOfPossibleMoves != 9 {
		t.Errorf("Was expecting at least 9 possible moves, but got %d", numberOfPossibleMoves)
	}
}

func TestBoard_MakingAMoveTakesPossibleMovesAway(t *testing.T) {
	anEmptyBoard := NewBoard()

	possibleMoves := anEmptyBoard.GetPossibleMoves()

	updatedBoard, _ := anEmptyBoard.MakeMove(possibleMoves[0], X)

	numberOfPossibleMoves := len(updatedBoard.GetPossibleMoves())

	if numberOfPossibleMoves != 8 {
		t.Errorf("Was expecting at least 8 possible moves, but got %d", numberOfPossibleMoves)
	}
}

func TestBoard_CantMarkTheSamePositionTwice(t *testing.T) {
	anEmptyBoard := NewBoard()

	updatedBoard, _ := anEmptyBoard.MakeMove(0, X)
	updatedBoard, err := updatedBoard.MakeMove(0, X)

	if err == nil {
		t.Errorf("Should not have been able to mark the same spot twice: %v", updatedBoard)
	}
}

func TestBoard_CantMarkOutOfBounds(t *testing.T) {
	anEmptyBoard := NewBoard()

	anEmptyBoard, err := anEmptyBoard.MakeMove(12, X)

	if err == nil {
		t.Errorf("Should not have been able to mark an impossible move on board: %v", anEmptyBoard)
	}
}
