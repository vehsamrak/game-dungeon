package app

type Room struct {
	x int
	y int
}

func (room *Room) X() int {
	return room.x
}

func (room *Room) Y() int {
	return room.y
}
