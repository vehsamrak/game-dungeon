package commands

import "github.com/vehsamrak/game-dungeon/internal/app"

type MoveCommand struct {
	roomRepository app.RoomRepository
}

func (command MoveCommand) Create(roomRepository app.RoomRepository) *MoveCommand {
	return &MoveCommand{roomRepository: roomRepository}
}

func (*MoveCommand) Name() string {
	return "move"
}

func (command *MoveCommand) Execute(character Character, arguments ...interface{}) (err error) {
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

	room := command.roomRepository.FindByXY(x, y)

	if room != nil && command.checkRoomMobility(room.Type()) {
		character.Move(x, y)
	}

	return
}

func (command *MoveCommand) checkRoomMobility(roomType string) bool {
	unmovableTypes := make(map[string]bool)
	unmovableTypes[app.RoomTypeMountain] = true

	return !unmovableTypes[roomType]
}
