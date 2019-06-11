package app

import (
	"fmt"
)

type GameMap struct {
	roomRepository RoomRepository
	fields         map[int]map[int]string
}

func (gameMap GameMap) Create(height int, width int, roomRepository RoomRepository) *GameMap {
	rooms := make(map[int]map[int]string)
	for x := 1; x <= height; x++ {
		roomsRow := make(map[int]string)
		for y := 1; y <= width; y++ {
			roomsRow[y] = fmt.Sprintf("%d.%d", x, y)
		}
		rooms[x] = roomsRow
	}

	return &GameMap{roomRepository: roomRepository, fields: rooms}
}

func (gameMap *GameMap) Size() (height int, width int) {
	return len(gameMap.fields), len(gameMap.fields[len(gameMap.fields)])
}

func (gameMap *GameMap) Room(x int, y int, z int) *Room {
	return gameMap.roomRepository.FindByXYandZ(x, y, z)
}
