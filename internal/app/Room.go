package app

import "github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"

type Room struct {
	x     int
	y     int
	flags map[roomFlag.Flag]bool
}

func (room Room) Create(x int, y int) *Room {
	return &Room{x: x, y: y, flags: make(map[roomFlag.Flag]bool)}
}

func (room *Room) X() int {
	return room.x
}

func (room *Room) Y() int {
	return room.y
}

func (room *Room) AddFlag(flag roomFlag.Flag) {
	room.flags[flag] = true
}

func (room *Room) AddFlags(flags []roomFlag.Flag) {
	for _, flag := range flags {
		room.flags[flag] = true
	}
}

func (room *Room) HasFlag(flag roomFlag.Flag) bool {
	return room.flags[flag]
}
