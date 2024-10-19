package main

import "time"

type GameStatus int

const (
	Initialize GameStatus = 0
	Play       GameStatus = 1
	Finished   GameStatus = 2
)

type GameInfo struct {
	score  int
	status GameStatus
	speed  int
	ticker *time.Ticker
}

func InitGameInfo() *GameInfo {
	return &GameInfo{
		score:  0,
		status: Initialize,
		speed:  300,
		ticker: time.NewTicker(time.Duration(300) * time.Millisecond),
	}
}

func (info *GameInfo) getScore() {
	info.score += 1
	if info.score%5 == 0 && info.speed > 50 {
		info.speed -= 20
		info.ticker.Stop()
		*info.ticker = *time.NewTicker(time.Duration(info.speed) * time.Millisecond)
	}
}
