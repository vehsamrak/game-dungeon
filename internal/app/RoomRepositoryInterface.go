package app

type RoomRepository interface {
	FindByXandY(x int, y int) *Room
	AddRoom(room *Room)
	FindByXY(XY XYInterface) *Room
}
