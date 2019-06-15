package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomBiom"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/random"
)

type MineCommand struct {
	roomRepository app.RoomRepository
	random         *random.Random
}

func (command *MineCommand) HealthPrice() int {
	return 5
}

func (command MineCommand) Create(roomRepository app.RoomRepository, random *random.Random) *MineCommand {
	return &MineCommand{roomRepository: roomRepository, random: random}
}

func (command *MineCommand) Execute(character Character, arguments ...string) (result CommandResult) {
	result = commandResult{}.Create()

	if !character.HasItemFlag(itemFlag.MineTool) {
		result.AddError(gameError.NoTool)

		return
	}

	room := command.roomRepository.FindByXYandZ(character.X(), character.Y(), character.Z())

	if room.Biom() != roomBiom.Mountain && room.Biom() != roomBiom.Cave {
		result.AddError(gameError.WrongBiom)

		return
	}

	if room.Biom() == roomBiom.Mountain {
		command.mineMountain(room, character, result)

		return
	}

	oreFound := command.random.RandomBoolean()
	if room != nil && room.HasFlag(roomFlag.OreProbability) && oreFound {
		ore := app.Item{}.Create()
		ore.AddFlag(itemFlag.ResourceOre)
		character.AddItem(ore)

		newCaveFound := command.random.RandomBoolean()
		if room.HasFlag(roomFlag.CaveProbability) && newCaveFound {
			newCaveDirections := []direction.Direction{
				direction.North,
				direction.South,
				direction.East,
				direction.West,
				direction.Down,
			}

			var caveError gameError.Error
			for i := len(newCaveDirections); i > 0; i-- {
				newCaveDirectionKey := command.random.RandomNumber(len(newCaveDirections) - 1)
				_, _, _, caveError = command.createCave(character, newCaveDirections[newCaveDirectionKey])

				if caveError == "" {
					break
				}

				newCaveDirections[newCaveDirectionKey] = newCaveDirections[len(newCaveDirections)-1]
				newCaveDirections = newCaveDirections[:len(newCaveDirections)-1]
			}

			if caveError != "" {
				result.AddError(caveError)
			}
		}

		room.RemoveFlag(roomFlag.CaveProbability)
	} else {
		result.AddError(gameError.OreNotFound)
		result.SetLowerHealthOnError(true)
	}

	return result
}

func (command *MineCommand) mineMountain(room *app.Room, character Character, result CommandResult) {
	if room.HasFlag(roomFlag.CaveProbability) {
		room.RemoveFlag(roomFlag.CaveProbability)

		if command.random.RandomBoolean() {
			result.AddError(gameError.CaveNotFound)

			return
		}

		x, y, z, err := command.createCave(character, direction.Down)

		if err == "" {
			character.Move(x, y, z)
		} else {
			result.AddError(err)
		}

	} else {
		result.AddError(gameError.CaveNotFound)
		result.SetLowerHealthOnError(true)
	}
}

func (command *MineCommand) createCave(
	character Character,
	newCaveDirection direction.Direction,
) (x int, y int, z int, err gameError.Error) {
	xDiff, yDiff, zDiff := newCaveDirection.DiffXYZ()
	x = character.X() + xDiff
	y = character.Y() + yDiff
	z = character.Z() + zDiff

	alreadyExistingRoom := command.roomRepository.FindByXYandZ(x, y, z)
	if alreadyExistingRoom != nil {
		err = gameError.RoomAlreadyExist
		return
	}

	newRoom := app.Room{}.Create(x, y, z, roomBiom.Cave)
	newRoom.AddFlag(roomFlag.OreProbability)

	if command.random.RandomBoolean() {
		newRoom.AddFlag(roomFlag.CaveProbability)
	}

	command.roomRepository.AddRoom(newRoom)

	return
}
