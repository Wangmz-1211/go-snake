package main

type Command interface {
	exec() bool
}

type baseCommand struct {
	mp    *GameMap
	snake *GameSnake
	food  *GameFood
}

func (base *baseCommand) move(dx, dy int) bool {
	mp := base.mp
	snake := base.snake
	food := base.food
	lastHeadPos := snake.getHeadPosition()
	nextHeadPos := Position{lastHeadPos.x + dx, lastHeadPos.y + dy}
	if !isValidPosition(nextHeadPos, mp) {
		// game over
		return false
	}
	eating := food.tryEat(nextHeadPos)
	if eating || food.timeLeft == 0 {
		food.gernerateNewFood(mp)
	}

	mp.SetCell(lastHeadPos, Body)
	mp.SetCell(nextHeadPos, Head)
	if eating {
		snake.eatAndMove(nextHeadPos)
		return true
	}

	lastTailPos := snake.getTailPosition()
	snake.moveForward(nextHeadPos)
	mp.SetCell(lastTailPos, Blank)
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
	return 0 <= p.x && p.x < m.Row && 0 <= p.y && p.y < m.Col
}
