package app

type RoomRepository interface {
	FindByXY(x int, y int) *Room
}