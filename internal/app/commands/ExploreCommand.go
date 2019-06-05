package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/exception"
)

type ExploreCommand struct {
	roomRepository app.RoomRepository
}

func (command ExploreCommand) Create(roomRepository app.RoomRepository) *ExploreCommand {
	return &ExploreCommand{roomRepository: roomRepository}
}

func (command *ExploreCommand) Execute(character Character, arguments ...interface{}) (err error) {
	exploreDirection := arguments[0].(direction.Direction)
	moveCommand := MoveCommand{}.Create(command.roomRepository)
	err = moveCommand.Execute(character, exploreDirection)

	if _, ok := err.(exception.RoomNotFound); ok {
		xDiff, yDiff := exploreDirection.DiffXY()
		x := character.X() + xDiff
		y := character.Y() + yDiff

		room := app.Room{}.Create(x, y)

		command.roomRepository.AddRoom(room)

		err = nil
	}

	return
}
