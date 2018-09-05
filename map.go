package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type maprender struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (m *maprender) New(w *ecs.World) {
	m.BasicEntity = ecs.NewBasic()
	m.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: 0,
			Y: 0,
		},
		Width:  worldMaxX,
		Height: worldMaxY,
	}

	common.CameraBounds = engo.AABB{
		Min: engo.Point{X: 0, Y: 0},
		Max: engo.Point{X: 0, Y: 0},
	}

	engo.Mailbox.Dispatch(common.CameraMessage{
		Axis:  common.YAxis,
		Value: m.SpaceComponent.Height / 2,
	})

	engo.Mailbox.Dispatch(common.CameraMessage{
		Axis:  common.XAxis,
		Value: m.SpaceComponent.Width / 2,
	})

	m.RenderComponent = common.RenderComponent{}

	texture, err := common.LoadedSprite("map.jpg")
	if err != nil {
		panic(err)
	}

	m.RenderComponent.Drawable = texture
	m.RenderComponent.Scale = engo.Point{X: 1, Y: 1}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&m.BasicEntity, &m.RenderComponent, &m.SpaceComponent)
		}
	}
}

func (m *maprender) Remove(e ecs.BasicEntity) {}

func (m *maprender) Update(dt float32) {
}
