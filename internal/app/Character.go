package app

type Character struct {
	name string
	x    int
	y    int
}

func (character *Character) Name() string {
	return character.name
}

func (character *Character) X() int {
	return character.x
}

func (character *Character) Y() int {
	return character.y
}

func (character *Character) Move(x int, y int) error {
	character.x = x
	character.y = y

	return nil
}
