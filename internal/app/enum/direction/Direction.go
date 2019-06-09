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

func (direction Direction) DiffXY() (x int, y int) {
	switch direction {
	case North:
		y += 1
	case South:
		y -= 1
	case East:
		x += 1
	case West:
		x -= 1
		// TODO[petr]: up, down, Z coordinates
	case Up:
		x += 100
		y += 100
	case Down:
		x -= 100
		y -= 100
	}

	return
}
