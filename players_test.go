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
		Entry("first row", 0, 1, 2),
		Entry("second row", 3, 4, 5),
		Entry("third row", 6, 7, 8),
		Entry("first column", 0, 3, 6),
		Entry("second column", 1, 4, 7),
		Entry("third column", 2, 5, 8),
		Entry("first diagonal", 0, 4, 8),
		Entry("second diagonal", 2, 4, 6),
	)
})
