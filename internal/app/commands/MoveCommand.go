package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomBiom"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/timer"
	"time"
)

type MoveCommand struct {
	roomRepository app.RoomRepository
	waitState      time.Duration
	healthPrice    int
}

func (command MoveCommand) Create(roomRepository app.RoomRepository) *MoveCommand {
	return &MoveCommand{
		roomRepository: roomRepository,
		waitState:      2 * time.Second,
		healthPrice:    1,
	}
}

func (command *MoveCommand) HealthPrice() int {
	return command.healthPrice
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

	if character.HasActiveTimers() {
		result.AddError(gameError.WaitState)
		return
	}

	character.SetTimer(timer.Move, command.waitState)

	xDiff, yDiff, zDiff := moveDirection.DiffXYZ()
	x := character.X() + xDiff
	y := character.Y() + yDiff
	z := character.Z() + zDiff

	initialRoom := command.roomRepository.FindByXYZ(character)
	destinationRoom := command.roomRepository.FindByXYandZ(x, y, z)

	err = command.checkRoomMobility(character, initialRoom, destinationRoom)
	if err == "" {
		character.Move(x, y, z)
	} else {
		result.AddError(err)
	}

	return
}

func (command *MoveCommand) checkRoomMobility(
	character Character,
	initialRoom *app.Room,
	destinationRoom *app.Room,
) (err gameError.Error) {
	if initialRoom == nil || destinationRoom == nil {
		err = gameError.RoomNotFound
	} else if initialRoom.Biom() == roomBiom.Water && destinationRoom.Biom() == roomBiom.Water {
		err = gameError.CantMoveInWater
	} else if destinationRoom.Biom() == roomBiom.Air && !character.HasItemFlag(itemFlag.CanFly) {
		err = gameError.RoomUnfordable
	} else if destinationRoom.Biom() == roomBiom.Cliff && !character.HasItemFlag(itemFlag.CliffWalk) {
		err = gameError.RoomUnfordable
	} else if destinationRoom.HasFlag(roomFlag.Unfordable) {
		err = gameError.RoomUnfordable
	}

	return
}
