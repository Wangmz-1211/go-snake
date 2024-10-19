package main

type Cell int

const (
	Blank Cell = 0
	Body  Cell = 1
	Head  Cell = 2
	Food  Cell = 3
)

func (c *Cell) getRune() rune {
	switch *c {
	case Blank:
		return '🐾'
	case Body:
		return '🚌'
	case Head:
		return '👶'
	case Food:
		return '🍎'
	}
	return '🐾'
}
