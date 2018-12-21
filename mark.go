package main

type Mark string

const (
	X Mark = "X"
	O Mark = "O"
)

func (mark Mark) Is(other *Mark) bool {
	return other != nil && mark == *other
}
