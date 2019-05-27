package game

type Game struct {
	Players []Player
}

type Player struct {
	X     int
	y     int
	Score int
}

