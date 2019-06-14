package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
)

type Character interface {
	Name() string
	X() int
	Y() int
	Z() int
	Move(x int, y int, z int)
	Inventory() []*app.Item
	AddItems(items []*app.Item)
	AddItem(item *app.Item)
	HasItemFlag(itemFlag itemFlag.Flag) bool
	Health() int
	MaxHealth() int
	LowerHealth(healthPoints int)
}
