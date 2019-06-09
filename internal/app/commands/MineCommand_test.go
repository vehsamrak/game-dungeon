package commands_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/notice"
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
	assert.True(suite.T(), character.HasItemFlag(itemFlag.MineTool))
	assert.True(suite.T(), commandResult.HasError(gameError.WrongBiom))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsCaveAndHasNoOreProbability_oreNotFound() {
	character := suite.createCharacterWithTool()
	roomRepository, room := suite.createRoomRepositoryWithCaveRoom(character, []roomFlag.Flag{})
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))

	commandResult := command.Execute(character)

	assert.Equal(suite.T(), roomBiom.Cave, room.Biom())
	assert.True(suite.T(), character.HasItemFlag(itemFlag.MineTool))
	assert.True(suite.T(), commandResult.HasError(gameError.OreNotFound))
}

func (suite *mineCommandTest) Test_Execute_characterWithToolAndRoomBiomIsCaveAndHasOreProbabilityAndNoOre_oreNotFound() {
	character := suite.createCharacterWithTool()
	roomRepository, room := suite.createRoomRepositoryWithCaveRoom(character, []roomFlag.Flag{roomFlag.OreProbability})
	command := commands.MineCommand{}.Create(roomRepository, suite.createRandomWithSeed(0))

	commandResult := command.Execute(character)

	assert.Equal(suite.T(), roomBiom.Cave, room.Biom())
	assert.True(suite.T(), character.HasItemFlag(itemFlag.MineTool))
	assert.True(suite.T(), commandResult.HasError(gameError.OreNotFound))
}

// TODO[petr]: Test_Execute_characterWithToolAndRoomBiomIsMountainAndNoCaveProbability_caveNotFoundCaveProbabilityRemoved
// TODO[petr]: Test_Execute_characterWithToolAndRoomBiomIsMountainAndCaveProbability_caveWithOreProbabilityAndWithCaveProbabilityOpenedDownAndCharacterMovedDown
// TODO[petr]: Test_Execute_characterWithToolAndRoomBiomIsMountainAndCaveProbability_caveWithOreProbabilityAndWithoutCaveProbabilityOpenedDownAndCharacterMovedDown
// TODO[petr]: Test_Execute_characterWithToolAndRoomBiomIsCaveAndHasOreProbabilityAndOreAndCaveProbability_oreFoundAndNewCaveOpened
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
	assert.False(suite.T(), commandResult.HasNotice(notice.FoundOre))
}

func (suite *mineCommandTest) createRoomRepositoryWithMountainRoom(
	character *app.Character,
	roomFlags []roomFlag.Flag,
) app.RoomRepository {
	room := app.Room{}.Create(character.X(), character.Y(), roomBiom.Mountain)
	room.AddFlags(roomFlags)

	return suite.createRoomRepositoryWithRoom(character.X(), character.Y(), roomBiom.Water, room)
}

func (suite *mineCommandTest) createRoomRepositoryWithCaveRoom(
	character *app.Character,
	roomFlags []roomFlag.Flag,
) (app.RoomRepository, *app.Room) {
	room := app.Room{}.Create(character.X(), character.Y(), roomBiom.Cave)
	room.AddFlags(roomFlags)

	return suite.createRoomRepositoryWithRoom(character.X(), character.Y(), roomBiom.Water, room), room
}

func (suite *mineCommandTest) createRoomRepositoryWithForestRoom(character *app.Character) (app.RoomRepository, *app.Room) {
	room := app.Room{}.Create(character.X(), character.Y(), roomBiom.Forest)

	return suite.createRoomRepositoryWithRoom(character.X(), character.Y(), roomBiom.Water, room), room
}

func (suite *mineCommandTest) createRoomRepositoryWithRoom(
	x int,
	y int,
	biom roomBiom.Biom,
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
