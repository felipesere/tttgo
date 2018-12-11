package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard_GetPossibleMoves(t *testing.T) {
	anEmptyBoard := NewBoard(9)

	numberOfPossibleMoves := len(anEmptyBoard.GetPossibleMoves())

	assert.Equal(t, 9, numberOfPossibleMoves)
}

func TestBoard_MakingAMoveTakesPossibleMovesAway(t *testing.T) {
	anEmptyBoard := NewBoard(9)

	possibleMoves := anEmptyBoard.GetPossibleMoves()

	updatedBoard, _ := anEmptyBoard.MakeMove(possibleMoves[0], X)

	numberOfPossibleMoves := len(updatedBoard.GetPossibleMoves())

	assert.Equal(t, 8, numberOfPossibleMoves)
}

func TestBoard_CantMarkTheSamePositionTwice(t *testing.T) {
	anEmptyBoard := NewBoard(9)

	updatedBoard, _ := anEmptyBoard.MakeMove(0, X)
	updatedBoard, err := updatedBoard.MakeMove(0, X)

	assert.NotNil(t, err, "Should not have been able to mark the same spot twice")
}

func TestBoard_CantMarkOutOfBounds(t *testing.T) {
	anEmptyBoard := NewBoard(9)

	anEmptyBoard, err := anEmptyBoard.MakeMove(12, X)

	assert.NotNil(t, err, "Should not have been able to mark an impossible move on board")
}

func TestBoard_FindsAWinnerInTheFirstColumn(t *testing.T) {
	anEmptyBoard := NewBoard(9)

	first, _ := anEmptyBoard.MakeMove(0, X)
	second, _ := first.MakeMove(1, X)
	third, _ := second.MakeMove(2, X)

	winner := third.Winner()

	assert.Equal(t, X, winner, "Was not able to find winner")
}
