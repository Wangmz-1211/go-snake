package main

type Command interface {
	exec() bool
}

type baseCommand struct {
	mp    *GameMap
	snake *GameSnake
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
	mp := cmd.base.mp
	snake := cmd.base.snake
	lastHeadPos := snake.getHeadPosition()
	nextHeadPos := Position{lastHeadPos.x - 1, lastHeadPos.y}
	if !isValidPosition(nextHeadPos, mp) {
		// game over
		return false
	}
	lastTailPos := snake.getTailPosition()
	snake.moveForward(nextHeadPos)
	mp.SetCell(lastHeadPos, Body)
	mp.SetCell(lastTailPos, Blank)
	mp.SetCell(nextHeadPos, Head)
	return true
}

func (cmd MoveRightCommand) exec() bool {
	mp := cmd.base.mp
	snake := cmd.base.snake
	lastHeadPos := snake.getHeadPosition()
	nextHeadPos := Position{lastHeadPos.x + 1, lastHeadPos.y}
	if !isValidPosition(nextHeadPos, mp) {
		// game over
		return false
	}
	lastTailPos := snake.getTailPosition()
	snake.moveForward(nextHeadPos)
	mp.SetCell(lastHeadPos, Body)
	mp.SetCell(lastTailPos, Blank)
	mp.SetCell(nextHeadPos, Head)
	return true
}

func (cmd MoveUpCommand) exec() bool {
	mp := cmd.base.mp
	snake := cmd.base.snake
	lastHeadPos := snake.getHeadPosition()
	nextHeadPos := Position{lastHeadPos.x, lastHeadPos.y - 1}
	if !isValidPosition(nextHeadPos, mp) {
		// game over
		return false
	}
	lastTailPos := snake.getTailPosition()
	snake.moveForward(nextHeadPos)
	mp.SetCell(lastHeadPos, Body)
	mp.SetCell(lastTailPos, Blank)
	mp.SetCell(nextHeadPos, Head)
	return true
}

func (cmd MoveDownCommand) exec() bool {
	mp := cmd.base.mp
	snake := cmd.base.snake
	lastHeadPos := snake.getHeadPosition()
	nextHeadPos := Position{lastHeadPos.x, lastHeadPos.y + 1}
	if !isValidPosition(nextHeadPos, mp) {
		// game over
		return false
	}
	lastTailPos := snake.getTailPosition()
	snake.moveForward(nextHeadPos)
	mp.SetCell(lastHeadPos, Body)
	mp.SetCell(lastTailPos, Blank)
	mp.SetCell(nextHeadPos, Head)
	return true
}

/* Utils */
func isValidPosition(p Position, m *GameMap) bool {
	return 0 <= p.x && p.x < m.Row && 0 <= p.y && p.y < m.Col
}
