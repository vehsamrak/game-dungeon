package direction

type Direction string

const (
	North Direction = "north"
	South Direction = "south"
	East  Direction = "east"
	West  Direction = "west"
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
	}

	return
}
