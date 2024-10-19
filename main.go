package main

import (
	"log"

	"github.com/nsf/termbox-go"
)

func main() {
	// Init termbox
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	// Alternative Screen
	termbox.HideCursor()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
	defer termbox.Flush()
	defer termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	defer termbox.SetCursor(1, 1)

	// Initialize Game
	row, col := termbox.Size()
	row /= 2
	info := InitGameInfo()
	mp := InitMap(row, col)
	snake := InitSnake(mp)
	food := InitFood(mp)

	key_binding := InitKeyBinding(info, mp, snake, food)
	eventCh := make(chan termbox.Event)
	go func() {
		for {
			event := termbox.PollEvent()
			eventCh <- event
		}
	}()
	var ev termbox.Event
	timer := info.ticker

	for {
		select {
		case <-timer.C:
			if ev.Key == termbox.KeyEsc {
				return
			}
			key_binding.invoke(ev.Key)
		case event := <-eventCh:
			ev = event
		}
	}
}
