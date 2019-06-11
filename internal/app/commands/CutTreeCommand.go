package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
)

type CutTreeCommand struct {
	roomRepository app.RoomRepository
}

func (command CutTreeCommand) Create(roomRepository app.RoomRepository) *CutTreeCommand {
	return &CutTreeCommand{roomRepository: roomRepository}
}

func (command *CutTreeCommand) Execute(character Character, arguments ...interface{}) (result CommandResult) {
	result = commandResult{}.Create()
	room := command.roomRepository.FindByXYandZ(character.X(), character.Y(), character.Z())

	if room != nil && character.HasItemFlag(itemFlag.CutTree) && room.HasFlag(roomFlag.Trees) {
		wood := app.Item{}.Create()
		wood.AddFlag(itemFlag.ResourceWood)
		character.AddItem(wood)
	} else {
		result.AddError(gameError.NoTool)
	}

	return
}
