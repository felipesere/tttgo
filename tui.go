package main

import (
	"fmt"
	"io"
	"strings"
)

type Display interface {
	show(board Board)
}

type Tui struct {
	writer io.Writer
}

func (tui Tui) show(board Board) {
	formatted := make([]string, board.size)

	for idx, row := range board.Rows() {
		formatted[idx] = formatRow(row)
	}
	io.WriteString(tui.writer, strings.Join(formatted, "\n"))
}

func formatRow(row Line) string {
	output := ""
	for _, cell := range row {
		if cell.empty {
			output += fmt.Sprintf("[%v]", cell.move+1)
		} else {
			output += fmt.Sprintf("[%v]", cell.mark)
		}
	}
	return output
}

func NewTui(output io.Writer) Display {
	return Tui{writer: output}
}
