package components

import ecs "github.com/x-hgg-x/goecs/v2"

type GameComponents struct {
	Player *ecs.NullComponent
}
type Components struct {
	Player *Player
}

type Player struct{}
