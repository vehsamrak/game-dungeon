package commands

type MoveCommand struct {
}

func (command MoveCommand) Create() *MoveCommand {
	return &MoveCommand{}
}

func (*MoveCommand) Name() string {
	return "move"
}

func (*MoveCommand) Execute(character Character, arguments ...interface{}) {
	direction := arguments[0]

	x := character.X()
	y := character.Y()

	switch direction {
	case "north":
		x -= 1
	case "south":
		x += 1
	case "east":
		y += 1
	case "west":
		y -= 1
	}

	character.Move(x, y)
}
