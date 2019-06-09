package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomBiom"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/random"
)

type MineCommand struct {
	roomRepository app.RoomRepository
	random         *random.Random
}

func (command MineCommand) Create(roomRepository app.RoomRepository, random *random.Random) *MineCommand {
	return &MineCommand{roomRepository: roomRepository, random: random}
}

func (command *MineCommand) Execute(character Character, arguments ...interface{}) (result CommandResult) {
	result = commandResult{}.Create()

	if !character.HasItemFlag(itemFlag.MineTool) {
		result.AddError(gameError.NoTool)

		return
	}

	room := command.roomRepository.FindByXY(character.X(), character.Y())

	if room.Biom() != roomBiom.Mountain && room.Biom() != roomBiom.Cave {
		result.AddError(gameError.WrongBiom)

		return
	}

	oreFound := command.random.RandomBoolean()
	if room != nil && room.HasFlag(roomFlag.OreProbability) && oreFound {
		ore := app.Item{}.Create()
		ore.AddFlag(itemFlag.ResourceOre)
		character.AddItem(ore)
	} else {
		result.AddError(gameError.OreNotFound)
	}

	return result
}
