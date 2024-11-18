package render

import (
	"game/components"
	"game/consts"

	ec "github.com/x-hgg-x/goecsengine/components"
	"github.com/x-hgg-x/goecsengine/loader"
	"github.com/x-hgg-x/goecsengine/math"
	w "github.com/x-hgg-x/goecsengine/world"
)

func RenderGrid(world w.World, screenPos math.Vector2, playerPos math.Vector2, gridState [][]int) {
	componentList := loader.EntityComponentList{}
	spriteSheet := (*world.Resources.SpriteSheets)["game"]
	for i := screenPos.Y; i < screenPos.Y+consts.GridHeight; i++ {
		for j := screenPos.X; j < screenPos.X+consts.GridWidth; j++ {
			componentList.Engine = append(componentList.Engine, loader.EngineComponentList{
				SpriteRender: &ec.SpriteRender{
					SpriteSheet:  &spriteSheet,
					SpriteNumber: 2,
				},
				Transform: ec.NewTransform().
					SetTranslation(
						float64(j*consts.TileSize+consts.TileSize/2),
						float64(i*consts.TileSize+consts.TileSize/2),
					).SetDepth(0),
			})
			componentList.Game = append(componentList.Game, components.Components{})
			if i == playerPos.Y && j == playerPos.X {
				componentList.Engine = append(componentList.Engine, loader.EngineComponentList{
					SpriteRender: &ec.SpriteRender{
						SpriteSheet:  &spriteSheet,
						SpriteNumber: 5,
					},
					Transform: ec.NewTransform().
						SetTranslation(
							float64(j*consts.TileSize+consts.TileSize/2),
							float64(i*consts.TileSize+consts.TileSize/2),
						).SetDepth(1),
				})
				componentList.Game = append(componentList.Game, components.Components{
					Player: &components.Player{},
				})
			}
		}
	}
	loader.AddEntities(world, componentList)
}
