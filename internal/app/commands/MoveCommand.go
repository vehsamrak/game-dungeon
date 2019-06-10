package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
)

type MoveCommand struct {
	roomRepository app.RoomRepository
}

func (command MoveCommand) Create(roomRepository app.RoomRepository) *MoveCommand {
	return &MoveCommand{roomRepository: roomRepository}
}

func (command *MoveCommand) Execute(character Character, arguments ...interface{}) (result CommandResult) {
	result = commandResult{}.Create()
	xDiff, yDiff := arguments[0].(direction.Direction).DiffXY()
	x := character.X() + xDiff
	y := character.Y() + yDiff

	room := command.roomRepository.FindByXandY(x, y)

	err := command.checkRoomMobility(room)

	if err == "" {
		character.Move(x, y)
	} else {
		result.AddError(err)
	}

	return
}

func (command *MoveCommand) checkRoomMobility(room *app.Room) (err gameError.Error) {
	if room == nil {
		err = gameError.RoomNotFound
	} else if room.HasFlag(roomFlag.Unfordable) {
		err = gameError.RoomUnfordable
	}

	return
}
