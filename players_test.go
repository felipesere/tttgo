package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("a computer player", func() {
	It("will defend an obvious position", func() {
		board := NewBoard(3)

		board, _ = board.MakeMove(0, X)
		board, _ = board.MakeMove(1, X)

		computer := NewComputer(O)

		defendedBoard := computer.MakeMove(board)

		Expect(defendedBoard.IsMovePossible(2)).To(Equal(false))
	})
})
