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
		y += 1
	case direction.South:
		y -= 1
	case direction.East:
		x += 1
	case direction.West:
		x -= 1
	}

	room := command.roomRepository.FindByXY(x, y)

	err = command.checkRoomMobility(room)

	if err == nil {
		character.Move(x, y)
	}

	return
}

func (command *MoveCommand) checkRoomMobility(room *app.Room) (err error) {
	if room == nil {
		err = exception.CantMove{}
	} else if room.HasFlag(app.RoomFlagUnfordable) {
		err = exception.RoomUnfordable{}
	}

	return
}
