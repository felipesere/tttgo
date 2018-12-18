package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
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

	DescribeTable("finds winning combinations",
		func(first, second, third int, winner Mark) {
			startingBoard := NewBoard(3)
			b, _ := startingBoard.MakeMove(first, winner)
			b, _ = b.MakeMove(second, winner)
			b, _ = b.MakeMove(third, winner)

			actualWinner := b.Winner()

			Expect(actualWinner).To(Equal(winner))
		},
		Entry("first row", 0, 1, 2, O),
		Entry("second row", 3, 4, 5, O),
		Entry("third row", 6, 7, 8, O),
		Entry("first column", 0, 3, 6, O),
		Entry("second column", 1, 4, 7, O),
		Entry("third column", 2, 5, 8, O),
		Entry("first diagonal", 0, 4, 8, O),
		Entry("second diagonal", 2, 4, 6, O),
	)
})
