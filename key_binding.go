package main

import (
	"github.com/nsf/termbox-go"
)

type KeyBinding struct {
	lastKey termbox.Key
	Map     *GameMap
	Snake   *GameSnake
	Mapping map[termbox.Key]Command
}

func InitKeyBinding(mp *GameMap, snake *GameSnake, food *GameFood) KeyBinding {
	mapping := make(map[termbox.Key]Command, 4)
	base := baseCommand{
		mp,
		snake,
		food,
	}
	mapping[termbox.KeyArrowLeft] = MoveLeftCommand{&base}
	mapping[termbox.KeyArrowDown] = MoveDownCommand{&base}
	mapping[termbox.KeyArrowRight] = MoveRightCommand{&base}
	mapping[termbox.KeyArrowUp] = MoveUpCommand{&base}
	return KeyBinding{
		lastKey: termbox.KeyArrowRight,
		Map:     mp,
		Snake:   snake,
		Mapping: mapping,
	}
}

func (k *KeyBinding) invoke(key termbox.Key) {
	if !k.isValidKey(key) {
		key = k.lastKey
	}
	if cmd := k.Mapping[key]; cmd != nil {
		if cmd.exec() {
			k.lastKey = key
		}
	}
}

func (k *KeyBinding) isValidKey(key termbox.Key) bool {
	switch key {
	case termbox.KeyArrowLeft:
		if k.lastKey != termbox.KeyArrowRight {
			return true
		}
	case termbox.KeyArrowRight:
		if k.lastKey != termbox.KeyArrowLeft {
			return true
		}
	case termbox.KeyArrowUp:
		if k.lastKey != termbox.KeyArrowDown {
			return true
		}
	case termbox.KeyArrowDown:
		if k.lastKey != termbox.KeyArrowUp {
			return true
		}
	default:
		return false
	}
	return false
}
