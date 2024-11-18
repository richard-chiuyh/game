package states

import (
	"fmt"
	"game/components"
	"game/consts"
	"game/render"

	ecs "github.com/x-hgg-x/goecs/v2"
	c "github.com/x-hgg-x/goecsengine/components"
	"github.com/x-hgg-x/goecsengine/math"
	"github.com/x-hgg-x/goecsengine/states"
	w "github.com/x-hgg-x/goecsengine/world"
)

type Pos struct{}

// GridState is the main game state
type GridState struct {
	screenPos math.Vector2
	playerPos math.Vector2
	gridState [][]int
}

// OnPause method
func (st *GridState) OnPause(world w.World) {}

// OnResume method
func (st *GridState) OnResume(world w.World) {}

// OnStart method
func (st *GridState) OnStart(world w.World) {
	render.RenderGrid(world, st.screenPos, st.playerPos, st.gridState)
}

// OnStop method
func (st *GridState) OnStop(world w.World) {
	world.Manager.DeleteAllEntities()
}

// Update method
func (st *GridState) Update(world w.World) states.Transition {
	gameComponents := world.Components.Game.(*components.GameComponents)
	moveUpAction := world.Resources.InputHandler.Actions["MoveUp"]
	moveDownAction := world.Resources.InputHandler.Actions["MoveDown"]
	moveLeftAction := world.Resources.InputHandler.Actions["MoveLeft"]
	moveRightAction := world.Resources.InputHandler.Actions["MoveRight"]
	switch {
	case moveUpAction:
		st.playerPos.Y += 1
	case moveDownAction:
		st.playerPos.Y -= 1
	case moveLeftAction:
		st.playerPos.X -= 1
	case moveRightAction:
		st.playerPos.X += 1
	}
	fmt.Println(st.playerPos.X, st.playerPos.Y)
	world.Manager.Join(gameComponents.Player, world.Components.Engine.Transform).Visit(ecs.Visit(func(entity ecs.Entity) {
		newPos := math.Vector2{
			X: float64(st.playerPos.X*consts.TileSize + consts.TileSize/2),
			Y: float64(st.playerPos.Y*consts.TileSize + consts.TileSize/2),
		}
		world.Components.Engine.Transform.Get(entity).(*c.Transform).Translation = newPos
	}))
	return states.Transition{}
}
