package direction

type Direction string

const (
	North Direction = "north"
	South Direction = "south"
	East  Direction = "east"
	West  Direction = "west"
	Up    Direction = "up"
	Down  Direction = "down"
)

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
