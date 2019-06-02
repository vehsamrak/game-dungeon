package app

type Room struct {
	x        int
	y        int
	roomType string
}

const RoomTypeRoad = "road"
const RoomTypeMountain = "mountain"
const RoomTypeForest = "forest"
const RoomTypeDeepForest = "deep_forest"

func (room Room) Create(x int, y int, roomType string) *Room {
	return &Room{x: x, y: y, roomType: roomType}
}

func (room *Room) X() int {
	return room.x
}

func (room *Room) Y() int {
	return room.y
}

func (room *Room) Type() string {
	return room.roomType
}
