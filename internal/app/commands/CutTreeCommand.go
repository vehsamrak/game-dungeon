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
	wood := app.Item{}.Create()
	wood.AddType(app.ItemTypeResourceWood)
	character.AddItem(wood)
	if room != nil && command.checkHasTools(character) && command.checkRoomType(room.Type()) {
	} else {
		err = errors.New("no tools or room has no trees")
	}

	return
}

func (command *CutTreeCommand) checkHasTools(character Character) bool {
	return character.HasType(app.ItemTypeCutTree)
}

func (command *CutTreeCommand) checkRoomType(roomType string) bool {
	typesWithTrees := map[string]bool{
		app.RoomTypeForest:     true,
		app.RoomTypeDeepForest: true,
	}

	return typesWithTrees[roomType]
}
