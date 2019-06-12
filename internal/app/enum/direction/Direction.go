package direction

import (
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
)

type Direction string

const (
	North Direction = "north"
	South Direction = "south"
	East  Direction = "east"
	West  Direction = "west"
	Up    Direction = "up"
	Down  Direction = "down"
)

func FromString(directionName string) (direction Direction, err gameError.Error) {
	directions := map[string]Direction{
		"north": North,
		"south": South,
		"east":  East,
		"west":  West,
		"up":    Up,
		"down":  Down,
	}

	direction, ok := directions[directionName]
	if !ok {
		err = gameError.WrongDirection
	}

	return
}

func (direction Direction) DiffXYZ() (x int, y int, z int) {
	switch direction {
	case North:
		y += 1
	case South:
		y -= 1
	case East:
		x += 1
	case West:
		x -= 1
	case Up:
		z += 1
	case Down:
		z -= 1
	}

	return
}

func (direction Direction) String() string {
	return string(direction)
}
