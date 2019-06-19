package mapGenerator

import "github.com/vehsamrak/game-dungeon/internal/app"

type WaterGenerator struct {
}

func (*WaterGenerator) generate() []*app.Room {
	return []*app.Room{}
}
