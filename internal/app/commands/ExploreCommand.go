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
	result = commandResult{}.Create()

	characterRoom := command.roomRepository.FindByXYZ(character)
	if characterRoom != nil && characterRoom.Biom() == roomBiom.Cave {
		result.AddError(gameError.WrongBiom)

		return
	}

	exploreDirection := arguments[0].(direction.Direction)
	moveCommand := MoveCommand{}.Create(command.roomRepository)
	result = moveCommand.Execute(character, exploreDirection)

	if result.HasError(gameError.RoomNotFound) {
		result.RemoveError(gameError.RoomNotFound)

		xDiff, yDiff, zDiff := exploreDirection.DiffXYZ()
		x := character.X() + xDiff
		y := character.Y() + yDiff
		z := character.Z() + zDiff

		biom := command.generateRandomBiom()
		room := app.Room{}.Create(x, y, z, biom)
		room.AddFlags(room.Biom().Flags())

		character.Move(x, y, z)

		command.roomRepository.AddRoom(room)
	}

	return
}

func (command *ExploreCommand) generateRandomBiom() roomBiom.Biom {
	bioms := roomBiom.All()
	randomNumber := command.random.RandomNumber(len(bioms) - 1)

	return bioms[randomNumber]
}
