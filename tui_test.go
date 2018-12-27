package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type CollectedOuput struct {
	printed string
}

func (co *CollectedOuput) Write(p []byte) (n int, err error) {
	co.printed += string(p)

	return len(p), nil
}

var _ = Describe("the TUI", func() {

	var (
		emptyBoard Board
		output     CollectedOuput
		tui        Display
	)

	BeforeEach(func() {
		emptyBoard = NewBoard(3)
		output = CollectedOuput{}
		tui = NewTui(&output)
	})

	It("prints a board", func() {
		tui.show(emptyBoard)

		Expect(output.printed).To(BeTheSameStringAs(
			`[1][2][3]
[4][5][6]
[7][8][9]`))
	})

	It("shows the placed markers", func() {
		board, _ := emptyBoard.MakeMove(3, X)
		board, _ = board.MakeMove(8, O)
		tui.show(board)

		Expect(output.printed).To(BeTheSameStringAs(
			`[1][2][X]
[4][5][6]
[7][O][9]`))
	})
})
