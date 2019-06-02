package commands

type Character interface {
	Name() string
	X() int
	Y() int
	Move(x int, y int) error
}
