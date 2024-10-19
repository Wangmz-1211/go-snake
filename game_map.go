package main

import "github.com/nsf/termbox-go"

type GameMap struct {
	Map [][]Cell
	Row int
	Col int
}

var GAME_MAP *GameMap

func InitMap(row, col int) *GameMap {
	if GAME_MAP == nil {
		// Create a instacne
		mp := make([][]Cell, row)
		for i := range mp {
			mp[i] = make([]Cell, col)
		}
		for i := range mp {
			for j := range mp[i] {
				termbox.SetCell(i*2, j, mp[i][j].getRune(), termbox.ColorDefault, termbox.ColorDefault)
			}
		}
		termbox.Flush()

		GAME_MAP = &GameMap{
			Map: mp,
			Row: row,
			Col: col,
		}
	}
	return GAME_MAP
}

func (m *GameMap) SetCell(p Position, target Cell) {
	m.Map[p.x][p.y] = target
	termbox.SetCell(p.x*2, p.y, target.getRune(), termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}
func (m *GameMap) GetCell(p Position) Cell {
	// unsafe
	return m.Map[p.x][p.y]
}

func (m *GameMap) isBlank(p Position) bool {
	return m.Map[p.x][p.y] == Blank
}
