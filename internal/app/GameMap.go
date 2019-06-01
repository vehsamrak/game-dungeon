package app

import (
	"fmt"
)

type GameMap struct {
	rooms map[int]map[int]string
}

func (gameMap GameMap) Create(height int, width int) *GameMap {
	rooms := make(map[int]map[int]string)
	for x := 1; x <= height; x++ {
		roomsRow := make(map[int]string)
		for y := 1; y <= width; y++ {
			roomsRow[y] = fmt.Sprintf("%d.%d", x, y)
		}
		rooms[x] = roomsRow
	}

	return &GameMap{rooms: rooms}
}

func (gameMap *GameMap) Size() (height int, width int) {
	return len(gameMap.rooms), len(gameMap.rooms[len(gameMap.rooms)])
}

func (gameMap *GameMap) RoomId(x int, y int) string {
	return gameMap.rooms[x][y]
}
