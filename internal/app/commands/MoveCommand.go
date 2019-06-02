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

	switch direction {
	case "north":
		character.Move(character.X()-1, character.Y())
	}
}
