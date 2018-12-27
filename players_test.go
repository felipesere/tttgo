package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("a computer player", func() {
	DescribeTable("how it defends in the",
		func(attack1, attack2, defense int) {
			board := NewBoard(3)

			board, _ = board.MakeMove(attack1, X)
			board, _ = board.MakeMove(attack2, X)

			computer := NewComputer(O)

			defendedBoard := computer.MakeMove(board)

			Expect(defendedBoard.IsMovePossible(defense)).To(Equal(false))
		},
		Entry("first row", 1, 2, 3),
		Entry("second row", 4, 5, 6),
		Entry("third row", 7, 8, 9),
		Entry("first column", 1, 4, 7),
		Entry("second column", 2, 5, 8),
		Entry("third column", 3, 6, 9),
		Entry("first diagonal", 1, 5, 9),
		Entry("second diagonal", 3, 5, 7),
	)
})
