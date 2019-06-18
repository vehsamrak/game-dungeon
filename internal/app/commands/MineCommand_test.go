package commands_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomBiom"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/random"
	"testing"
)

func TestMineCommand(test *testing.T) {
	suite.Run(test, &mineCommandTest{})
}

type mineCommandTest struct {
	suite.Suite
}

func (suite *mineCommandTest) Test_Execute_characterWithoutTool_noTool() {
	character := suite.createCharacterWithoutTool()
	roomRepository := &app.RoomMemoryRepository{}
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))

	commandResult := command.Execute(character)

	assert.False(suite.T(), character.HasItemFlag(itemFlag.MineTool))
	assert.True(suite.T(), commandResult.HasError(gameError.NoTool))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsNotMountainOrCave_wrongBiom() {
	character := suite.createCharacterWithTool()
	roomRepository, room := suite.createRoomRepositoryWithForestRoom(character)
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))

	commandResult := command.Execute(character)

	assert.Equal(suite.T(), roomBiom.Forest, room.Biom())
	assert.True(suite.T(), commandResult.HasError(gameError.WrongBiom))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsMountainAndNoCaveProbability_caveNotFound() {
	character := suite.createCharacterWithTool()
	roomRepository, _ := suite.createRoomRepositoryWithMountainRoom(character, []roomFlag.Flag{})
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(0))

	commandResult := command.Execute(character)

	assert.True(suite.T(), commandResult.HasError(gameError.CaveNotFound))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsMountainAndCaveProbability_caveProbabilityRemovedAndCaveWithOreProbabilityAndWithCaveProbabilityOpenedDownAndCharacterMovedDown() {
	character := suite.createCharacterWithTool()
	characterBeforeCommand := *character
	roomRepository, initialRoom := suite.createRoomRepositoryWithMountainRoom(character, []roomFlag.Flag{roomFlag.CaveProbability})
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(3))
	targetX, targetY, targetZ := direction.Down.DiffXYZ()

	commandResult := command.Execute(character)

	newRoom := roomRepository.FindByXYandZ(targetX, targetY, targetZ)
	assert.False(suite.T(), commandResult.HasErrors())
	assert.Equal(suite.T(), roomBiom.Cave, newRoom.Biom())
	assert.True(suite.T(), newRoom.HasFlag(roomFlag.OreProbability))
	assert.True(suite.T(), newRoom.HasFlag(roomFlag.CaveProbability))
	assert.NotEqual(
		suite.T(),
		fmt.Sprintf("%d%d%d", characterBeforeCommand.X(), characterBeforeCommand.Y(), characterBeforeCommand.Z()),
		fmt.Sprintf("%d%d%d", character.X(), character.Y(), character.Z()),
	)
	assert.Equal(suite.T(), targetX, character.X())
	assert.Equal(suite.T(), targetY, character.Y())
	assert.Equal(suite.T(), targetZ, character.Z())
	assert.False(suite.T(), initialRoom.HasFlag(roomFlag.CaveProbability))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsMountainAndCaveProbability_caveProbabilityRemovedAndCaveWithOreProbabilityAndWithoutCaveProbabilityOpenedDownAndCharacterMovedDown() {
	character := suite.createCharacterWithTool()
	characterBeforeCommand := *character
	roomRepository, initialRoom := suite.createRoomRepositoryWithMountainRoom(character, []roomFlag.Flag{roomFlag.CaveProbability})
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(0))
	targetX, targetY, targetZ := direction.Down.DiffXYZ()

	commandResult := command.Execute(character)

	newRoom := roomRepository.FindByXYandZ(targetX, targetY, targetZ)
	assert.False(suite.T(), commandResult.HasErrors())
	assert.False(suite.T(), initialRoom.HasFlag(roomFlag.CaveProbability))
	assert.Equal(suite.T(), roomBiom.Cave, newRoom.Biom())
	assert.True(suite.T(), newRoom.HasFlag(roomFlag.OreProbability))
	assert.False(suite.T(), newRoom.HasFlag(roomFlag.CaveProbability))
	assert.NotEqual(
		suite.T(),
		fmt.Sprintf("%d%d%d", characterBeforeCommand.X(), characterBeforeCommand.Y(), characterBeforeCommand.Z()),
		fmt.Sprintf("%d%d%d", character.X(), character.Y(), character.Z()),
	)
	assert.Equal(suite.T(), targetX, character.X())
	assert.Equal(suite.T(), targetY, character.Y())
	assert.Equal(suite.T(), targetZ, character.Z())
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsMountainAndCaveProbability_caveProbabilityRemovedAndCaveNotFoundAndCharacterNotMoved() {
	character := suite.createCharacterWithTool()
	characterBeforeCommand := *character
	roomRepository, room := suite.createRoomRepositoryWithMountainRoom(character, []roomFlag.Flag{roomFlag.CaveProbability})
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))
	targetX, targetY, targetZ := direction.Down.DiffXYZ()

	commandResult := command.Execute(character)

	assert.False(suite.T(), room.HasFlag(roomFlag.CaveProbability))
	assert.True(suite.T(), commandResult.HasError(gameError.CaveNotFound))
	assert.Equal(suite.T(), characterBeforeCommand.X()+characterBeforeCommand.Y(), character.X()+character.Y())
	assert.Nil(suite.T(), roomRepository.FindByXYandZ(targetX, targetY, targetZ))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsCaveAndHasNoOreProbability_oreNotFound() {
	character := suite.createCharacterWithTool()
	roomRepository, room := suite.createRoomRepositoryWithCaveRoom(character, []roomFlag.Flag{})
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))

	commandResult := command.Execute(character)

	assert.Equal(suite.T(), roomBiom.Cave, room.Biom())
	assert.True(suite.T(), commandResult.HasError(gameError.OreNotFound))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsCaveAndHasOreProbabilityAndNoOre_oreNotFound() {
	character := suite.createCharacterWithTool()
	roomRepository, _ := suite.createRoomRepositoryWithCaveRoom(character, []roomFlag.Flag{roomFlag.OreProbability})
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(0))

	commandResult := command.Execute(character)

	assert.True(suite.T(), commandResult.HasError(gameError.OreNotFound))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsCaveAndHasOreProbabilityAndOreAndCaveProbabilityAndNoNearRooms_orePlacedToCharacterInventoryAndNewCaveOpenedAndCaveProbabilityRemovedFromInitialRoom() {
	character := suite.createCharacterWithTool()
	roomRepository, room := suite.createRoomRepositoryWithCaveRoom(
		character,
		[]roomFlag.Flag{roomFlag.OreProbability, roomFlag.CaveProbability},
	)
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))

	result := command.Execute(character)
	secondResult := command.Execute(character)

	assert.False(suite.T(), result.HasErrors())
	assert.True(suite.T(), secondResult.HasError(gameError.WaitState))
	assert.True(suite.T(), character.HasItemFlag(itemFlag.ResourceOre))
	assert.True(suite.T(), suite.isNearCaveOpened(roomRepository))
	assert.False(suite.T(), room.HasFlag(roomFlag.CaveProbability))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsCaveAndHasOreProbabilityAndOreAndCaveProbabilityAndAllNearRoomsAlreadyExist_orePlacedToCharacterInventoryAndCaveNotOpened() {
	character := suite.createCharacterWithTool()
	northX, northY, northZ := direction.North.DiffXYZ()
	southX, southY, southZ := direction.South.DiffXYZ()
	eastX, eastY, eastZ := direction.East.DiffXYZ()
	westX, westY, westZ := direction.West.DiffXYZ()
	downX, downY, downZ := direction.Down.DiffXYZ()
	initialRoom := app.Room{}.Create(character.X(), character.Y(), character.Z(), roomBiom.Cave)
	initialRoom.AddFlags([]roomFlag.Flag{roomFlag.OreProbability, roomFlag.CaveProbability})
	rooms := []*app.Room{
		initialRoom,
		app.Room{}.Create(northX, northY, northZ, roomBiom.Forest),
		app.Room{}.Create(southX, southY, southZ, roomBiom.Forest),
		app.Room{}.Create(eastX, eastY, eastZ, roomBiom.Forest),
		app.Room{}.Create(westX, westY, westZ, roomBiom.Forest),
		app.Room{}.Create(downX, downY, downZ, roomBiom.Forest),
	}
	roomRepository := suite.createRoomRepositoryWithRooms(rooms)
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))

	commandResult := command.Execute(character)

	assert.True(suite.T(), commandResult.HasError(gameError.RoomAlreadyExist))
	assert.True(suite.T(), character.HasItemFlag(itemFlag.ResourceOre))
	assert.False(suite.T(), suite.isNearCaveOpened(roomRepository))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsCaveAndHasOreProbabilityAndOreAndCaveProbabilityAndTargetRoomAlreadyExist_orePlacedToCharacterInventoryAndNewCaveOpenedInAnotherDirection() {
	character := suite.createCharacterWithTool()
	southX, southY, southZ := direction.South.DiffXYZ()
	eastX, eastY, eastZ := direction.East.DiffXYZ()
	westX, westY, westZ := direction.West.DiffXYZ()
	initialRoom := app.Room{}.Create(character.X(), character.Y(), character.Z(), roomBiom.Cave)
	initialRoom.AddFlags([]roomFlag.Flag{roomFlag.OreProbability, roomFlag.CaveProbability})
	rooms := []*app.Room{
		initialRoom,
		app.Room{}.Create(southX, southY, southZ, roomBiom.Forest),
		app.Room{}.Create(eastX, eastY, eastZ, roomBiom.Forest),
		app.Room{}.Create(westX, westY, westZ, roomBiom.Forest),
	}
	roomRepository := suite.createRoomRepositoryWithRooms(rooms)
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))

	commandResult := command.Execute(character)

	assert.False(suite.T(), commandResult.HasErrors())
	assert.True(suite.T(), character.HasItemFlag(itemFlag.ResourceOre))
	assert.True(suite.T(), suite.isNearCaveOpened(roomRepository))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsCaveAndHasOreProbabilityAndOreAndNoCaveProbability_orePlacedToCharacterInventory() {
	character := suite.createCharacterWithTool()
	roomRepository, _ := suite.createRoomRepositoryWithCaveRoom(
		character,
		[]roomFlag.Flag{roomFlag.OreProbability},
	)
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))

	commandResult := command.Execute(character)

	assert.False(suite.T(), commandResult.HasErrors())
	assert.True(suite.T(), character.HasItemFlag(itemFlag.ResourceOre))
	assert.False(suite.T(), suite.isNearCaveOpened(roomRepository))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsCaveAndHasOreProbabilityAndOre_orePlacedToCharacterInventory() {
	character := suite.createCharacterWithTool()
	characterBeforeCommand := *character
	roomRepository, room := suite.createRoomRepositoryWithCaveRoom(character, []roomFlag.Flag{roomFlag.OreProbability})
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))

	commandResult := command.Execute(character)

	assert.False(suite.T(), commandResult.HasErrors())
	assert.Equal(suite.T(), roomBiom.Cave, room.Biom())
	assert.True(suite.T(), character.HasItemFlag(itemFlag.MineTool))
	assert.False(suite.T(), characterBeforeCommand.HasItemFlag(itemFlag.ResourceOre))
	assert.True(suite.T(), character.HasItemFlag(itemFlag.ResourceOre))
}

func (suite *mineCommandTest) createRoomRepositoryWithMountainRoom(
	character *app.Character,
	roomFlags []roomFlag.Flag,
) (app.RoomRepository, *app.Room) {
	room := app.Room{}.Create(character.X(), character.Y(), character.Z(), roomBiom.Mountain)
	room.AddFlags(roomFlags)

	return suite.createRoomRepositoryWithRooms([]*app.Room{room}), room
}

func (suite *mineCommandTest) createRoomRepositoryWithCaveRoom(
	character *app.Character,
	roomFlags []roomFlag.Flag,
) (app.RoomRepository, *app.Room) {
	room := app.Room{}.Create(character.X(), character.Y(), character.Z(), roomBiom.Cave)
	room.AddFlags(roomFlags)

	return suite.createRoomRepositoryWithRooms([]*app.Room{room}), room
}

func (suite *mineCommandTest) createRoomRepositoryWithForestRoom(character *app.Character) (app.RoomRepository, *app.Room) {
	room := app.Room{}.Create(character.X(), character.Y(), character.Z(), roomBiom.Forest)

	return suite.createRoomRepositoryWithRooms([]*app.Room{room}), room
}

func (suite *mineCommandTest) createRoomRepositoryWithRooms(rooms []*app.Room) app.RoomRepository {
	return app.RoomMemoryRepository{}.Create(rooms)
}

func (suite *mineCommandTest) createCharacterWithoutTool() *app.Character {
	character := &app.Character{}

	return character
}

func (suite *mineCommandTest) createCharacterWithTool() *app.Character {
	tool := app.Item{}.Create("")
	tool.AddFlag(itemFlag.MineTool)
	character := app.Character{}.Create("")
	character.AddItem(tool)

	return character
}

func (suite *mineCommandTest) createRandomWithSeed(seed int64) *random.Randomizer {
	randomizer := random.Randomizer{}.Create()
	randomizer.Seed(seed)

	return randomizer
}

func (suite *mineCommandTest) isNearCaveOpened(roomRepository app.RoomRepository) (nearCaveOpened bool) {
	directions := []direction.Direction{
		direction.North,
		direction.South,
		direction.East,
		direction.West,
		direction.Down,
	}

	for _, searchDirection := range directions {
		x, y, z := searchDirection.DiffXYZ()
		newCave := roomRepository.FindByXYandZ(x, y, z)
		if newCave != nil && newCave.Biom() == roomBiom.Cave {
			return true
		}
	}

	return
}
