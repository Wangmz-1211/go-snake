package main

import "math/rand/v2"

type Position struct {
	x int
	y int
}

func generateRandomPosition(row, col int) Position {
	return Position{
		x: rand.IntN(row),
		y: rand.IntN(col),
	}
}
