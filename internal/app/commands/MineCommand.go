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

func (command MineCommand) Create(roomRepository app.RoomRepository, random *random.Random) *MineCommand {
	return &MineCommand{roomRepository: roomRepository, random: random}
}

func (command *MineCommand) Execute(character Character, arguments ...interface{}) (result CommandResult) {
	result = commandResult{}.Create()

	if !character.HasItemFlag(itemFlag.MineTool) {
		result.AddError(gameError.NoTool)

		return
	}

	room := command.roomRepository.FindByXY(character.X(), character.Y())

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
			newCaveDirectionKey := command.random.RandomNumber(3)
			newCaveDirections := []direction.Direction{
				direction.North,
				direction.South,
				direction.East,
				direction.West,
			}

			command.createCave(character, newCaveDirections[newCaveDirectionKey])
		}
	} else {
		result.AddError(gameError.OreNotFound)
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

		x, y := command.createCave(character, direction.Down)

		character.Move(x, y)
	} else {
		result.AddError(gameError.CaveNotFound)
	}
}

func (command *MineCommand) createCave(character Character, newCaveDirection direction.Direction) (x int, y int) {
	xDiff, yDiff := newCaveDirection.DiffXY()
	x = character.X() + xDiff
	y = character.Y() + yDiff

	newRoom := app.Room{}.Create(x, y, roomBiom.Cave)
	newRoom.AddFlag(roomFlag.OreProbability)

	if command.random.RandomBoolean() {
		newRoom.AddFlag(roomFlag.CaveProbability)
	}

	command.roomRepository.AddRoom(newRoom)

	return x, y
}
