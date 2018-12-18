package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("a board", func() {

	var (
		anEmptyBoard Board
	)

	BeforeEach(func() {
		anEmptyBoard = NewBoard(3)
	})

	It("has 9 possible moves", func() {
		numberOfPossibleMoves := len(anEmptyBoard.GetPossibleMoves())
		Expect(numberOfPossibleMoves).To(Equal(9))
	})

	It("has fewer possible moves when moves are made", func() {
		possibleMoves := anEmptyBoard.GetPossibleMoves()

		updatedBoard, _ := anEmptyBoard.MakeMove(possibleMoves[0], X)

		numberOfPossibleMoves := len(updatedBoard.GetPossibleMoves())

		Expect(numberOfPossibleMoves).To(Equal(8))
	})

	It("cant mark the same position twice", func() {
		updatedBoard, _ := anEmptyBoard.MakeMove(0, X)
		updatedBoard, err := updatedBoard.MakeMove(0, X)

		Expect(err).To(HaveOccurred())
	})

	It("cant mark outside of the bounds of the board", func() {
		_, err := anEmptyBoard.MakeMove(12, X)

		Expect(err).To(HaveOccurred())
	})

	It("finds a winner in the first row", func() {
		first, _ := anEmptyBoard.MakeMove(0, O)
		second, _ := first.MakeMove(1, O)
		third, _ := second.MakeMove(2, O)

		winner := third.Winner()

		Expect(winner).To(Equal(O))
	})

	It("finds a winner in the first column", func() {
		first, _ := anEmptyBoard.MakeMove(0, X)
		second, _ := first.MakeMove(3, X)
		third, _ := second.MakeMove(6, X)

		winner := third.Winner()

		Expect(winner).To(Equal(X))
	})

	It("finds a winner in the first diagonal", func() {
		first, _ := anEmptyBoard.MakeMove(0, X)
		second, _ := first.MakeMove(4, X)
		third, _ := second.MakeMove(8, X)

		winner := third.Winner()

		Expect(winner).To(Equal(X))
	})

	It("finds a winner in the second diagonal", func() {
		first, _ := anEmptyBoard.MakeMove(2, X)
		second, _ := first.MakeMove(4, X)
		third, _ := second.MakeMove(6, X)

		winner := third.Winner()

		Expect(winner).To(Equal(X))
	})
})
