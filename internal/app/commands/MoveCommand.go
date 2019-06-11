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
	xDiff, yDiff, zDiff := arguments[0].(direction.Direction).DiffXYZ()
	x := character.X() + xDiff
	y := character.Y() + yDiff
	z := character.Z() + zDiff

	room := command.roomRepository.FindByXYandZ(x, y, z)

	err := command.checkRoomMobility(room)

	if err == "" {
		character.Move(x, y, z)
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
