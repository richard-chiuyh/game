package main

import (
	"game/components"
	"game/consts"
	"game/controls"
	"game/states"

	"github.com/x-hgg-x/goecsengine/loader"
	r "github.com/x-hgg-x/goecsengine/resources"
	s "github.com/x-hgg-x/goecsengine/states"
	"github.com/x-hgg-x/goecsengine/utils"
	w "github.com/x-hgg-x/goecsengine/world"

	"github.com/hajimehoshi/ebiten/v2"
)

type mainGame struct {
	world        w.World
	stateMachine s.StateMachine
}

func (game *mainGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return consts.WindowWidth, consts.WindowHeight
}

func (game *mainGame) Update() error {
	game.stateMachine.Update(game.world)
	return nil
}

func (game *mainGame) Draw(screen *ebiten.Image) {
	game.stateMachine.Draw(game.world, screen)
}

func main() {
	world := w.InitWorld(&components.GameComponents{})

	// Init screen dimensions
	world.Resources.ScreenDimensions = &r.ScreenDimensions{Width: consts.WindowWidth, Height: consts.WindowHeight}

	// Init controls
	axes := []string{}
	actions := []string{
		controls.MoveUpAction, controls.MoveDownAction, controls.MoveLeftAction, controls.MoveRightAction,
	}
	controls, inputHandler := loader.LoadControls("etc/controls.toml", axes, actions)
	world.Resources.Controls = &controls
	world.Resources.InputHandler = &inputHandler

	// Init fonts
	fonts := loader.LoadFonts("etc/fonts.toml")
	world.Resources.Fonts = &fonts

	// Init sprite sheets
	spriteSheets := loader.LoadSpriteSheets("etc/spritesheets.toml")
	world.Resources.SpriteSheets = &spriteSheets

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(consts.WindowWidth, consts.WindowHeight)
	ebiten.SetWindowTitle("")

	utils.LogError(ebiten.RunGame(&mainGame{world, s.Init(&states.GridState{}, world)}))
}
