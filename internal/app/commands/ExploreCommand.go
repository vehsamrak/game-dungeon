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

func (command *ExploreCommand) HealthPrice() int {
	return 3
}

func (command ExploreCommand) Create(roomRepository app.RoomRepository, random *random.Random) *ExploreCommand {
	return &ExploreCommand{roomRepository: roomRepository, random: random}
}

func (command *ExploreCommand) Execute(character Character, arguments ...string) (result CommandResult) {
	result = commandResult{}.Create()

	if len(arguments) < 1 {
		result.AddError(gameError.WrongCommandAttributes)
		return
	}

	exploreDirection, err := direction.FromString(arguments[0])
	if err != "" {
		result.AddError(gameError.WrongCommandAttributes)
		return
	}

	directionAllowed := command.checkDirection(exploreDirection)
	if !directionAllowed {
		result.AddError(gameError.WrongDirection)
		return
	}

	err = command.checkInitialRoom(command.roomRepository.FindByXYZ(character))
	if err != "" {
		result.AddError(err)
		return
	}

	moveCommand := MoveCommand{}.Create(command.roomRepository)
	result = moveCommand.Execute(character, exploreDirection.String())

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
	var biomProbabilities []roomBiom.Biom
	for _, biom := range roomBiom.All() {
		for i := 0; i < biom.ExploreProbability(); i++ {
			biomProbabilities = append(biomProbabilities, biom)
		}
	}

	randomNumber := command.random.RandomNumber(len(biomProbabilities) - 1)

	return biomProbabilities[randomNumber]
}

func (command *ExploreCommand) checkInitialRoom(room *app.Room) (err gameError.Error) {
	disallowedBioms := map[roomBiom.Biom]bool{
		roomBiom.Water: true,
		roomBiom.Cave:  true,
		roomBiom.Cliff: true,
		roomBiom.Air:   true,
	}

	if room == nil {
		err = gameError.RoomNotFound
	} else if _, biomIsDissalowed := disallowedBioms[room.Biom()]; biomIsDissalowed {
		err = gameError.WrongBiom
	}

	return
}

func (command *ExploreCommand) checkDirection(exploreDirection direction.Direction) (directionAllowed bool) {
	allowedDirections := map[direction.Direction]bool{
		direction.North: true,
		direction.South: true,
		direction.West:  true,
		direction.East:  true,
	}

	_, directionAllowed = allowedDirections[exploreDirection]

	return
}
