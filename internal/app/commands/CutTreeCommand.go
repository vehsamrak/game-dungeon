package commands

import (
	"errors"
	"github.com/vehsamrak/game-dungeon/internal/app"
)

type CutTreeCommand struct {
	roomRepository app.RoomRepository
}

func (command CutTreeCommand) Create(roomRepository app.RoomRepository) *CutTreeCommand {
	return &CutTreeCommand{roomRepository: roomRepository}
}

func (*CutTreeCommand) Name() string {
	return "cut trees"
}

func (command *CutTreeCommand) Execute(character Character, arguments ...interface{}) (err error) {
	room := command.roomRepository.FindByXY(character.X(), character.Y())

	if room != nil && character.HasItemFlag(app.ItemFlagCutTree) && room.HasFlag(app.RoomFlagHasTrees) {
		wood := app.Item{}.Create()
		wood.AddFlag(app.ItemFlagResourceWood)

		character.AddItem(wood)
	} else {
		err = errors.New("no tools or room has no trees")
	}

	return
}

func (command *CutTreeCommand) checkHasTools(character Character) bool {
	return character.HasItemFlag(app.ItemFlagCutTree)
}

func (command *CutTreeCommand) checkRoomFlags(room *app.Room) bool {
	return room.HasFlag(app.RoomFlagHasTrees)
}
