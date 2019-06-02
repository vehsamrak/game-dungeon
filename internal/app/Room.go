package app

type Room struct {
	x         int
	y         int
	roomTypes map[string]bool
}

const RoomTypeRoad = "road"
const RoomTypeUnfordable = "unfordable"
const RoomTypeForest = "forest"
const RoomTypeDeepForest = "deep_forest"

func (room Room) Create(x int, y int) *Room {
	return &Room{x: x, y: y, roomTypes: make(map[string]bool)}
}

func (room *Room) X() int {
	return room.x
}

func (room *Room) Y() int {
	return room.y
}

func (room *Room) AddType(roomType string) {
	room.roomTypes[roomType] = true
}

func (room *Room) AddTypes(roomTypes []string) {
	for _, roomType := range roomTypes {
		room.roomTypes[roomType] = true
	}
}

func (room *Room) HasType(itemType string) bool {
	return room.roomTypes[itemType]
}
