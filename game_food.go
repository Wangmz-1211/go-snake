package main

type GameFood struct {
	position Position
	timeLeft int
}

func InitFood(mp *GameMap) *GameFood {
	var pos Position
	for {
		pos = generateRandomPosition(mp.Row, mp.Col)
		if mp.isBlank(pos) {
			break
		}
	}
	mp.SetCell(pos, Food)
	return &GameFood{
		position: pos,
		timeLeft: 2 * (mp.Row + mp.Col),
	}
}

func (f *GameFood) gernerateNewFood(mp *GameMap) {
	var pos Position
	for {
		pos = generateRandomPosition(mp.Row, mp.Col)
		if mp.isBlank(pos) {
			break
		}
	}
	f.position = pos
	f.timeLeft = 2 * (mp.Row + mp.Col)
	mp.SetCell(pos, Food)
}

func (f *GameFood) tryEat(pos Position) bool {
	if pos.x == f.position.x && pos.y == f.position.y {
		return true
	}
	f.timeLeft -= 1
	return false
}
