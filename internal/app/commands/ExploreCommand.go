package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomBiom"
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

	if result.HasError(gameError.RoomNotFound) {
		result.RemoveError(gameError.RoomNotFound)

		xDiff, yDiff := exploreDirection.DiffXY()
		x := character.X() + xDiff
		y := character.Y() + yDiff

		biom := command.generateRandomBiom()
		room := app.Room{}.Create(x, y, biom)
		room.AddFlags(room.Biom().Flags())

		character.Move(x, y)

		command.roomRepository.AddRoom(room)
	}

	return
}

func (command *ExploreCommand) generateRandomBiom() roomBiom.Biom {
	bioms := roomBiom.All()
	randomNumber := command.random.RandomNumber(len(bioms) - 1)

	return bioms[randomNumber]
}
