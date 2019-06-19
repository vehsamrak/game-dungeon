package mapGenerator

import "github.com/vehsamrak/game-dungeon/internal/app"

type MapGenerator interface {
	generate() []*app.Room
}
