package commands_test

import (
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

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsCaveAndHasOreProbabilityAndNoOre_oreNotFound() {
	character := suite.createCharacterWithTool()
	roomRepository, _ := suite.createRoomRepositoryWithCaveRoom(character, []roomFlag.Flag{roomFlag.OreProbability})
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(0))

	commandResult := command.Execute(character)

	assert.True(suite.T(), commandResult.HasError(gameError.OreNotFound))
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
	targetX, targetY := direction.Down.DiffXY()

	commandResult := command.Execute(character)

	newRoom := roomRepository.FindByXY(targetX, targetY)
	assert.False(suite.T(), commandResult.HasErrors())
	assert.Equal(suite.T(), roomBiom.Cave, newRoom.Biom())
	assert.True(suite.T(), newRoom.HasFlag(roomFlag.OreProbability))
	assert.True(suite.T(), newRoom.HasFlag(roomFlag.CaveProbability))
	assert.NotEqual(suite.T(), characterBeforeCommand.X()+characterBeforeCommand.Y(), character.X()+character.Y())
	assert.Equal(suite.T(), targetX, character.X())
	assert.Equal(suite.T(), targetY, character.Y())
	assert.False(suite.T(), initialRoom.HasFlag(roomFlag.CaveProbability))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsMountainAndCaveProbability_caveProbabilityRemovedAndCaveWithOreProbabilityAndWithoutCaveProbabilityOpenedDownAndCharacterMovedDown() {
	character := suite.createCharacterWithTool()
	characterBeforeCommand := *character
	roomRepository, initialRoom := suite.createRoomRepositoryWithMountainRoom(character, []roomFlag.Flag{roomFlag.CaveProbability})
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(0))
	targetX, targetY := direction.Down.DiffXY()

	commandResult := command.Execute(character)

	newRoom := roomRepository.FindByXY(targetX, targetY)
	assert.False(suite.T(), commandResult.HasErrors())
	assert.False(suite.T(), initialRoom.HasFlag(roomFlag.CaveProbability))
	assert.Equal(suite.T(), roomBiom.Cave, newRoom.Biom())
	assert.True(suite.T(), newRoom.HasFlag(roomFlag.OreProbability))
	assert.False(suite.T(), newRoom.HasFlag(roomFlag.CaveProbability))
	assert.NotEqual(suite.T(), characterBeforeCommand.X()+characterBeforeCommand.Y(), character.X()+character.Y())
	assert.Equal(suite.T(), targetX, character.X())
	assert.Equal(suite.T(), targetY, character.Y())
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsMountainAndCaveProbability_caveProbabilityRemovedAndCaveNotFoundAndCharacterNotMoved() {
	character := suite.createCharacterWithTool()
	characterBeforeCommand := *character
	roomRepository, room := suite.createRoomRepositoryWithMountainRoom(character, []roomFlag.Flag{roomFlag.CaveProbability})
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))
	targetX, targetY := direction.Down.DiffXY()

	commandResult := command.Execute(character)

	assert.False(suite.T(), room.HasFlag(roomFlag.CaveProbability))
	assert.True(suite.T(), commandResult.HasError(gameError.CaveNotFound))
	assert.Equal(suite.T(), characterBeforeCommand.X()+characterBeforeCommand.Y(), character.X()+character.Y())
	assert.Nil(suite.T(), roomRepository.FindByXY(targetX, targetY))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsCaveAndHasNoOreProbability_oreNotFound() {
	character := suite.createCharacterWithTool()
	roomRepository, room := suite.createRoomRepositoryWithCaveRoom(character, []roomFlag.Flag{})
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))

	commandResult := command.Execute(character)

	assert.Equal(suite.T(), roomBiom.Cave, room.Biom())
	assert.True(suite.T(), commandResult.HasError(gameError.OreNotFound))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsCaveAndHasOreProbabilityAndOreAndCaveProbability_orePlacedToCharacterInventoryAndNewCaveOpened() {
	character := suite.createCharacterWithTool()
	roomRepository, _ := suite.createRoomRepositoryWithCaveRoom(
		character,
		[]roomFlag.Flag{roomFlag.OreProbability, roomFlag.CaveProbability},
	)
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))

	commandResult := command.Execute(character)

	assert.False(suite.T(), commandResult.HasErrors())
	suite.assertNearCaveOpened(roomRepository)
	assert.True(suite.T(), character.HasItemFlag(itemFlag.ResourceOre))
}

// TODO[petr]: Test_Execute_characterWithToolAndRoomBiomIsCaveAndHasOreProbabilityAndOreAndNoCaveProbability_orePlacedToCharacterInventory
// TODO[petr]: Test_Execute_characterWithToolAndRoomBiomIsCaveAndHasOreProbabilityAndNoOre_oreNotFound

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
	room := app.Room{}.Create(character.X(), character.Y(), roomBiom.Mountain)
	room.AddFlags(roomFlags)

	return suite.createRoomRepositoryWithRoom(character.X(), character.Y(), room), room
}

func (suite *mineCommandTest) createRoomRepositoryWithCaveRoom(
	character *app.Character,
	roomFlags []roomFlag.Flag,
) (app.RoomRepository, *app.Room) {
	room := app.Room{}.Create(character.X(), character.Y(), roomBiom.Cave)
	room.AddFlags(roomFlags)

	return suite.createRoomRepositoryWithRoom(character.X(), character.Y(), room), room
}

func (suite *mineCommandTest) createRoomRepositoryWithForestRoom(character *app.Character) (app.RoomRepository, *app.Room) {
	room := app.Room{}.Create(character.X(), character.Y(), roomBiom.Forest)

	return suite.createRoomRepositoryWithRoom(character.X(), character.Y(), room), room
}

func (suite *mineCommandTest) createRoomRepositoryWithRoom(
	x int,
	y int,
	room *app.Room,
) app.RoomRepository {
	return app.RoomMemoryRepository{}.Create([]*app.Room{room})
}

func (suite *mineCommandTest) createCharacterWithoutTool() *app.Character {
	character := &app.Character{}

	return character
}

func (suite *mineCommandTest) createCharacterWithTool() *app.Character {
	tool := app.Item{}.Create()
	tool.AddFlag(itemFlag.MineTool)
	character := &app.Character{}
	character.AddItem(tool)

	return character
}

func (suite *mineCommandTest) createRandomWithSeed(seed int64) *random.Random {
	randomizer := random.Random{}.Create()
	randomizer.Seed(seed)

	return randomizer
}

func (suite *mineCommandTest) assertNearCaveOpened(roomRepository app.RoomRepository) {
	directions := []direction.Direction{direction.North, direction.South, direction.East, direction.West}

	var nearCaveOpened bool
	for _, searchDirection := range directions {
		x, y := searchDirection.DiffXY()
		newCave := roomRepository.FindByXY(x, y)
		if newCave != nil && newCave.Biom() == roomBiom.Cave {
			nearCaveOpened = true
		}
	}

	assert.True(suite.T(), nearCaveOpened)
}
