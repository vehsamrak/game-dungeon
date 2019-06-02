package commands

import "github.com/vehsamrak/game-dungeon/internal/app"

type Character interface {
	Name() string
	X() int
	Y() int
	Move(x int, y int) error
	Inventory() []*app.Item
	AddItems(items []*app.Item)
	AddItem(item *app.Item)
	HasItemFlag(itemFlag string) bool
}
