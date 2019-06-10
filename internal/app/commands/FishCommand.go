package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomBiom"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/random"
)

type FishCommand struct {
	roomRepository app.RoomRepository
	random         *random.Random
}

func (command FishCommand) Create(roomRepository app.RoomRepository, random *random.Random) *FishCommand {
	return &FishCommand{roomRepository: roomRepository, random: random}
}

func (command *FishCommand) Execute(character Character, arguments ...interface{}) (result CommandResult) {
	result = commandResult{}.Create()

	if !character.HasItemFlag(itemFlag.FishTool) {
		result.AddError(gameError.NoTool)

		return
	}

	room := command.roomRepository.FindByXandY(character.X(), character.Y())

	if room.Biom() != roomBiom.Water {
		result.AddError(gameError.WrongBiom)

		return
	}

	resourceFound := command.random.RandomBoolean()
	if room != nil && room.HasFlag(roomFlag.FishProbability) && resourceFound {
		fish := app.Item{}.Create()
		fish.AddFlag(itemFlag.ResourceFish)
		character.AddItem(fish)
	} else {
		result.AddError(gameError.FishNotFound)
	}

	return result
}
