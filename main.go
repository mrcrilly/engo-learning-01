package main

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const (
	worldMaxX = 3565
	worldMaxY = 2519
)

type mapscene struct {
	world *ecs.World
}

func (s *mapscene) Type() string { return "Map Renderer" }

func (s *mapscene) Preload() {
	err := engo.Files.Load("map.jpg")
	if err != nil {
		panic(err)
	}
}

func (s *mapscene) Setup(u engo.Updater) {
	s.world = u.(*ecs.World)

	common.SetBackground(color.Black)
	common.MinZoom = 0.5
	common.MaxZoom = 2

	s.world.AddSystem(&common.RenderSystem{})
	s.world.AddSystem(common.NewKeyboardScroller(750, engo.DefaultHorizontalAxis, engo.DefaultVerticalAxis))
	s.world.AddSystem(&common.MouseZoomer{ZoomSpeed: -0.125})

	s.world.AddSystem(&maprender{})
}

func main() {
	config := engo.RunOptions{
		Title:      "Map Renderer",
		Fullscreen: true,
		// Width:          worldMaxX,
		// Height:         worldMaxY,
		VSync:          true,
		StandardInputs: true,
	}

	engo.Run(config, &mapscene{})
}
