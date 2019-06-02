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

	if room != nil && command.checkHasTools(character) && command.checkRoomType(room.Type()) {
	} else {
		err = errors.New("no tools or room has no trees")
	}

	return
}

func (command *CutTreeCommand) checkHasTools(character Character) bool {
	for _, item := range character.Inventory() {
		if item.HasType(app.ItemTypeCutTree) {
			return true
		}
	}

	return false
}

func (command *CutTreeCommand) checkRoomType(roomType string) bool {
	typesWithTrees := map[string]bool{
		app.RoomTypeForest:     true,
		app.RoomTypeDeepForest: true,
	}

	return typesWithTrees[roomType]
}
