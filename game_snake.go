package main

type GameSnake struct {
	Bodies []Position
	Ate    bool
}

func InitSnake(mp *GameMap) *GameSnake {
	row := mp.Row
	col := mp.Col

	row /= 2
	col /= 2

	bodies := make([]Position, 3)

	for i := range 3 {
		bodies[i] = Position{row - i, col}
		if i == 0 {
			mp.SetCell(bodies[i], Head)
		} else {
			mp.SetCell(bodies[i], Body)
		}
	}

	return &GameSnake{
		Bodies: bodies,
		Ate:    false,
	}

}

func (s *GameSnake) getHeadPosition() Position {
	return s.Bodies[0]
}

func (s *GameSnake) getTailPosition() Position {
	return s.Bodies[len(s.Bodies)-1]
}

func (s *GameSnake) moveForward(p Position) {
	s.Bodies = append([]Position{p}, s.Bodies[:len(s.Bodies)-1]...)
}

func (s *GameSnake) eatAndMove(p Position) {
	s.Bodies = append([]Position{p}, s.Bodies...)
}
