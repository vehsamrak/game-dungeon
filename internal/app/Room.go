package app

type Room struct {
	x     int
	y     int
	flags map[string]bool
}

const RoomFlagRoad = "road"
const RoomFlagUnfordable = "unfordable"
const RoomFlagTrees = "trees"

func (room Room) Create(x int, y int) *Room {
	return &Room{x: x, y: y, flags: make(map[string]bool)}
}

func (room *Room) X() int {
	return room.x
}

func (room *Room) Y() int {
	return room.y
}

func (room *Room) AddFlag(flag string) {
	room.flags[flag] = true
}

func (room *Room) AddFlags(flags []string) {
	for _, flag := range flags {
		room.flags[flag] = true
	}
}

func (room *Room) HasFlag(flag string) bool {
	return room.flags[flag]
}
