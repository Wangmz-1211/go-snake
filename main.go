package main

import (
	"log"
	"time"

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
	mp := InitMap(row, col)
	snake := InitSnake(mp)
	food := InitFood(mp)

	key_binding := InitKeyBinding(mp, snake, food)
	eventCh := make(chan termbox.Event)
	go func() {
		for {
			event := termbox.PollEvent()
			eventCh <- event
		}
	}()
	gap := 300
	var ev termbox.Event
	timer := time.NewTicker(time.Duration(gap) * time.Millisecond)

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
