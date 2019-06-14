package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomBiom"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
)

type MoveCommand struct {
	roomRepository app.RoomRepository
}

func (command MoveCommand) Create(roomRepository app.RoomRepository) *MoveCommand {
	return &MoveCommand{roomRepository: roomRepository}
}

func (command *MoveCommand) HealthPrice() int {
	return 1
}

func (command *MoveCommand) Execute(character Character, arguments ...string) (result CommandResult) {
	result = commandResult{}.Create()

	if len(arguments) < 1 {
		result.AddError(gameError.WrongCommandAttributes)
		return
	}

	moveDirection, err := direction.FromString(arguments[0])
	if err != "" {
		result.AddError(gameError.WrongCommandAttributes)
		return
	}

	xDiff, yDiff, zDiff := moveDirection.DiffXYZ()
	x := character.X() + xDiff
	y := character.Y() + yDiff
	z := character.Z() + zDiff

	initialRoom := command.roomRepository.FindByXYZ(character)
	destinationRoom := command.roomRepository.FindByXYandZ(x, y, z)

	err = command.checkRoomMobility(initialRoom, destinationRoom)
	if err == "" {
		character.Move(x, y, z)
	} else {
		result.AddError(err)
	}

	return
}

func (command *MoveCommand) checkRoomMobility(initialRoom *app.Room, destinationRoom *app.Room) (err gameError.Error) {
	if initialRoom == nil || destinationRoom == nil {
		err = gameError.RoomNotFound
	} else if initialRoom.Biom() == roomBiom.Water && destinationRoom.Biom() == roomBiom.Water {
		err = gameError.CantMoveFromWater
	} else if destinationRoom.HasFlag(roomFlag.Unfordable) {
		err = gameError.RoomUnfordable
	}

	return
}
