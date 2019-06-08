package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/exception"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/random"
)

type ExploreCommand struct {
	roomRepository app.RoomRepository
	random         *random.Random
}

func (command ExploreCommand) Create(roomRepository app.RoomRepository, random *random.Random) *ExploreCommand {
	return &ExploreCommand{roomRepository: roomRepository, random: random}
}

func (command *ExploreCommand) Execute(character Character, arguments ...interface{}) (result CommandResult) {
	exploreDirection := arguments[0].(direction.Direction)
	moveCommand := MoveCommand{}.Create(command.roomRepository)
	result = moveCommand.Execute(character, exploreDirection)

	if result.HasError(exception.RoomNotFound{}) {
		result.RemoveError(exception.RoomNotFound{})

		xDiff, yDiff := exploreDirection.DiffXY()
		x := character.X() + xDiff
		y := character.Y() + yDiff

		room := app.Room{}.Create(x, y)
		command.generateFlags(room)

		character.Move(x, y)

		command.roomRepository.AddRoom(room)
	}

	return
}

func (command *ExploreCommand) generateFlags(room *app.Room) {
	biomFlags := roomFlag.BiomFlags()
	randomNumber := command.random.RandomNumber(len(biomFlags) - 1)

	flag := biomFlags[randomNumber]

	room.AddFlag(flag)
}
