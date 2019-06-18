package app

import "github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"

type Item struct {
	name  string
	flags map[itemFlag.Flag]bool
}

func (item Item) Create(name string) *Item {
	return &Item{name: name, flags: make(map[itemFlag.Flag]bool)}
}

func (item *Item) Name() string {
	return item.name
}

func (item *Item) AddFlag(flag itemFlag.Flag) {
	item.flags[flag] = true
}

func (item *Item) HasFlag(flag itemFlag.Flag) bool {
	return item.flags[flag]
}
