package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/exception"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
)

type MoveCommand struct {
	roomRepository app.RoomRepository
}

func (command MoveCommand) Create(roomRepository app.RoomRepository) *MoveCommand {
	return &MoveCommand{roomRepository: roomRepository}
}

func (command *MoveCommand) Execute(character Character, arguments ...interface{}) (err error) {
	xDiff, yDiff := arguments[0].(direction.Direction).DiffXY()
	x := character.X() + xDiff
	y := character.Y() + yDiff

	room := command.roomRepository.FindByXY(x, y)

	err = command.checkRoomMobility(room)

	if err == nil {
		character.Move(x, y)
	}

	return
}

func (command *MoveCommand) checkRoomMobility(room *app.Room) (err error) {
	if room == nil {
		err = exception.RoomNotFound{}
	} else if room.HasFlag(roomFlag.Unfordable) {
		err = exception.RoomUnfordable{}
	}

	return
}
