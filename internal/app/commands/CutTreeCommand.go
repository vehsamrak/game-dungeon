package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/exception"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
)

type CutTreeCommand struct {
	roomRepository app.RoomRepository
}

func (command CutTreeCommand) Create(roomRepository app.RoomRepository) *CutTreeCommand {
	return &CutTreeCommand{roomRepository: roomRepository}
}

func (command *CutTreeCommand) Execute(character Character, arguments ...interface{}) (err error) {
	room := command.roomRepository.FindByXY(character.X(), character.Y())

	if room != nil && character.HasItemFlag(app.ItemFlagCutTree) && room.HasFlag(roomFlag.Trees) {
		wood := app.Item{}.Create()
		wood.AddFlag(app.ItemFlagResourceWood)

		character.AddItem(wood)
	} else {
		err = exception.NoTool{}
	}

	return
}
