package app

type RoomRepository interface {
	AddRoom(room *Room)
	FindByXYandZ(x int, y int, z int) *Room
	FindByXYZ(XYZ XYInterface) *Room
}
