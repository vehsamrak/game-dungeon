package app

type RoomRepository interface {
	Create() RoomRepository
	FindByXY(x int, y int) *Room
}
