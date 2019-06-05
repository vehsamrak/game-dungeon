package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/exception"
)

type MoveCommand struct {
	roomRepository app.RoomRepository
}

func (command MoveCommand) Create(roomRepository app.RoomRepository) *MoveCommand {
	return &MoveCommand{roomRepository: roomRepository}
}

func (command *MoveCommand) Execute(character Character, arguments ...interface{}) (err error) {
	x := character.X()
	y := character.Y()

	switch arguments[0] {
	case direction.North:
		x -= 1
	case direction.South:
		x += 1
	case direction.East:
		y += 1
	case direction.West:
		y -= 1
	}

	room := command.roomRepository.FindByXY(x, y)

	if room != nil && command.checkRoomMobility(room) {
		character.Move(x, y)
	} else {
		err = &exception.CantMove{}
	}

	return
}

func (command *MoveCommand) checkRoomMobility(room *app.Room) bool {
	return !room.HasFlag(app.RoomFlagUnfordable)
}
