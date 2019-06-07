package app

import "github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"

type Item struct {
	flags map[itemFlag.Flag]bool
}

func (item Item) Create() *Item {
	return &Item{flags: make(map[itemFlag.Flag]bool)}
}

func (item *Item) AddFlag(flag itemFlag.Flag) {
	item.flags[flag] = true
}

func (item *Item) HasFlag(flag itemFlag.Flag) bool {
	return item.flags[flag]
}
