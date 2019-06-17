package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomBiom"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/timer"
	"github.com/vehsamrak/game-dungeon/internal/app/random"
	"time"
)

type ExploreCommand struct {
	roomRepository app.RoomRepository
	random         *random.Randomizer
	waitState      time.Duration
	healthPrice    int
}

func (command *ExploreCommand) HealthPrice() int {
	return command.healthPrice
}

func (command ExploreCommand) Create(roomRepository app.RoomRepository, random *random.Randomizer) *ExploreCommand {
	return &ExploreCommand{
		roomRepository: roomRepository,
		random:         random,
		waitState:      3 * time.Second,
		healthPrice:    3,
	}
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

		if character.TimerActive(timer.Explore) {
			result.AddError(gameError.WaitState)
			return
		}

		character.SetTimer(timer.Explore, command.waitState)

		xDiff, yDiff, zDiff := exploreDirection.DiffXYZ()
		x := character.X() + xDiff
		y := character.Y() + yDiff
		z := character.Z() + zDiff

		bioms := command.applicableBioms(character, x, y, z)
		biom := command.selectRandomBiom(bioms)

		room := app.Room{}.Create(x, y, z, biom)
		room.AddFlags(biom.Flags())

		if command.isBiomMovable(biom, character) {
			character.Move(x, y, z)
		}

		command.roomRepository.AddRoom(room)
	}

	return
}

func (command *ExploreCommand) isBiomMovable(biom roomBiom.Biom, character Character) bool {
	airBiomIsMovable := biom != roomBiom.Air || biom == roomBiom.Air && character.HasItemFlag(itemFlag.CanFly)
	cliffBiomIsMovable := biom != roomBiom.Cliff || biom == roomBiom.Cliff && character.HasItemFlag(itemFlag.CliffWalk)

	return airBiomIsMovable && cliffBiomIsMovable
}

func (command *ExploreCommand) applicableBioms(
	character Character,
	x int,
	y int,
	z int,
) (applicableBioms []roomBiom.Biom) {
	bioms := roomBiom.All()
	if character.Z() <= 0 {
		for i, biom := range bioms {
			if biom == roomBiom.Air {
				applicableBioms = append(bioms[:i], bioms[i+1:]...)
			}
		}
	} else if character.Z() > 0 {
		for _, biom := range bioms {
			allowedBioms := map[roomBiom.Biom]bool{
				roomBiom.Air:      true,
				roomBiom.Mountain: true,
				roomBiom.Cliff:    true,
			}

			_, biomIsAllowed := allowedBioms[biom]
			if biomIsAllowed {
				bottomRoom := command.roomRepository.FindByXYandZ(x, y, z-1)
				if biom == roomBiom.Air || bottomRoom != nil && bottomRoom.Biom() == roomBiom.Mountain {
					applicableBioms = append(applicableBioms, biom)
				}
			}
		}
	}

	return applicableBioms
}

func (command *ExploreCommand) selectRandomBiom(bioms []roomBiom.Biom) roomBiom.Biom {
	var biomProbabilities []roomBiom.Biom
	for _, biom := range bioms {
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
