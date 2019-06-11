package app

import (
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomBiom"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
)

type Room struct {
	x     int
	y     int
	z     int
	biom  roomBiom.Biom
	flags map[roomFlag.Flag]bool
}

func (room Room) Create(x int, y int, z int, biom roomBiom.Biom) *Room {
	return &Room{x: x, y: y, z: z, biom: biom, flags: make(map[roomFlag.Flag]bool)}
}

func (room *Room) X() int {
	return room.x
}

func (room *Room) Y() int {
	return room.y
}

func (room *Room) Z() int {
	return room.z
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

func (room *Room) Biom() roomBiom.Biom {
	return room.biom
}

func (room *Room) Flags() map[roomFlag.Flag]bool {
	return room.flags
}

func (room *Room) RemoveFlag(flag roomFlag.Flag) {
	delete(room.flags, flag)
}
