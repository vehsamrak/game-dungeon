package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
)

type ExploreCommand struct {
	roomRepository app.RoomRepository
}

func (command ExploreCommand) Create(roomRepository app.RoomRepository) *ExploreCommand {
	return &ExploreCommand{roomRepository: roomRepository}
}

func (command *ExploreCommand) Execute(character Character, arguments ...interface{}) (err error) {
	return
}
