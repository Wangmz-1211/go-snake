package main

type Command interface {
	exec() bool
}

type baseCommand struct {
	info  *GameInfo
	mp    *GameMap
	snake *GameSnake
	food  *GameFood
}

func (base *baseCommand) move(dx, dy int) bool {
	lastHeadPos := base.snake.getHeadPosition()
	nextHeadPos := Position{lastHeadPos.x + dx, lastHeadPos.y + dy}
	if !isValidPosition(nextHeadPos, base.mp) {
		// game over
		base.info.status = Finished
		return false
	}
	eating := base.food.tryEat(nextHeadPos)
	if eating || base.food.timeLeft == 0 {
		base.food.gernerateNewFood(base.mp)
	}

	base.mp.SetCell(lastHeadPos, Body)
	base.mp.SetCell(nextHeadPos, Head)
	if eating {
		base.info.getScore()
		base.snake.eatAndMove(nextHeadPos)
		return true
	}

	lastTailPos := base.snake.getTailPosition()
	base.snake.moveForward(nextHeadPos)
	base.mp.SetCell(lastTailPos, Blank)
	return true
}

/* Command Structs */
type MoveLeftCommand struct {
	base *baseCommand
}

type MoveRightCommand struct {
	base *baseCommand
}

type MoveUpCommand struct {
	base *baseCommand
}

type MoveDownCommand struct {
	base *baseCommand
}

/* Interface Implements */
func (cmd MoveLeftCommand) exec() bool {
	return cmd.base.move(-1, 0)
}

func (cmd MoveRightCommand) exec() bool {
	return cmd.base.move(1, 0)
}

func (cmd MoveUpCommand) exec() bool {
	return cmd.base.move(0, -1)
}

func (cmd MoveDownCommand) exec() bool {
	return cmd.base.move(0, 1)
}

/* Utils */
func isValidPosition(p Position, m *GameMap) bool {
	return 0 <= p.x && p.x < m.Row && 0 <= p.y && p.y < m.Col && (m.isBlank(p) || m.isFood(p))
}
