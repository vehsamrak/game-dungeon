package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/exception"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/notice"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/random"
)

type SearchOreCommand struct {
	roomRepository app.RoomRepository
	random         *random.Random
}

func (command SearchOreCommand) Create(roomRepository app.RoomRepository, random *random.Random) *SearchOreCommand {
	return &SearchOreCommand{roomRepository: roomRepository, random: random}
}

func (command *SearchOreCommand) Execute(character Character, arguments ...interface{}) (result CommandResult) {
	result = commandResult{}.Create()

	if !character.HasItemFlag(itemFlag.SearchOre) {
		result.AddError(exception.NoTool{})

		return
	}

	room := command.roomRepository.FindByXY(character.X(), character.Y())

	oreFound := command.random.RandomBoolean()
	if room != nil && room.HasFlag(roomFlag.OreProbability) && oreFound {
		result.AddNotice(notice.FoundOre)
	} else {
		result.AddError(exception.OreNotFound{})
	}

	return result
}
