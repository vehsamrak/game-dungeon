package commands_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomBiom"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/random"
	"testing"
)

func TestFishCommand(test *testing.T) {
	suite.Run(test, &fishCommandTest{})
}

type fishCommandTest struct {
	suite.Suite
}

func (suite *fishCommandTest) Test_Execute_characterWithoutTool_noTool() {
	character := suite.createCharacterWithoutTools()
	roomRepository := &app.RoomMemoryRepository{}
	command := commands.FishCommand{}.Create(roomRepository, suite.createRandomWithSeed(0))

	commandResult := command.Execute(character)

	assert.True(suite.T(), commandResult.HasError(gameError.NoTool))
}

func (suite *fishCommandTest) Test_Execute_characterWithToolAndRoomIsNotWaterBiom_wrongBiomError() {
	character := suite.createCharacterWithTool()
	roomRepository := suite.createRoomRepositoryWithForestRoom(character)
	command := commands.FishCommand{}.Create(roomRepository, suite.createRandomWithSeed(0))

	commandResult := command.Execute(character)

	assert.True(suite.T(), commandResult.HasError(gameError.WrongBiom))
}

func (suite *fishCommandTest) Test_Execute_characterWithToolAndRoomIsWaterBiomAndHasNoFishProbability_fishNotFound() {
	character := suite.createCharacterWithTool()
	roomRepository := suite.createRoomRepositoryWithWaterRoom(character, []roomFlag.Flag{})
	command := commands.FishCommand{}.Create(roomRepository, suite.createRandomWithSeed(0))

	commandResult := command.Execute(character)

	assert.True(suite.T(), commandResult.HasError(gameError.FishNotFound))
}

func (suite *fishCommandTest) Test_Execute_characterWithToolAndRoomIsWaterBiomAndHasFishProbabilityAndRandomRolledNoFish_fishNotFound() {
	character := suite.createCharacterWithTool()
	characterBeforeCommand := *character
	roomRepository := suite.createRoomRepositoryWithWaterRoom(character, []roomFlag.Flag{roomFlag.FishProbability})
	command := commands.FishCommand{}.Create(roomRepository, suite.createRandomWithSeed(0))

	commandResult := command.Execute(character)

	assert.True(suite.T(), commandResult.HasError(gameError.FishNotFound))
	assert.Equal(suite.T(), character.Inventory(), characterBeforeCommand.Inventory())
	assert.False(suite.T(), character.HasItemFlag(itemFlag.ResourceFish))
}

func (suite *fishCommandTest) Test_Execute_characterWithToolAndRoomIsWaterBiomAndHasFishProbabilityAndRandomRolledFish_fishAppearsInCharacterInventory() {
	character := suite.createCharacterWithTool()
	characterBeforeCommand := *character
	roomRepository := suite.createRoomRepositoryWithWaterRoom(character, []roomFlag.Flag{roomFlag.FishProbability})
	command := commands.FishCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))

	commandResult := command.Execute(character)

	assert.False(suite.T(), commandResult.HasErrors())
	assert.NotEqual(suite.T(), character.Inventory(), characterBeforeCommand.Inventory())
	assert.True(suite.T(), character.HasItemFlag(itemFlag.ResourceFish))
}

func (suite *fishCommandTest) createRoomRepositoryWithWaterRoom(
	character *app.Character,
	roomFlags []roomFlag.Flag,
) app.RoomRepository {
	room := app.Room{}.Create(character.X(), character.Y(), roomBiom.Water)
	room.AddFlags(roomFlags)

	return suite.createRoomRepositoryWithRoom(character.X(), character.Y(), room)
}

func (suite *fishCommandTest) createRoomRepositoryWithForestRoom(character *app.Character) app.RoomRepository {
	room := app.Room{}.Create(character.X(), character.Y(), roomBiom.Forest)

	return suite.createRoomRepositoryWithRoom(character.X(), character.Y(), room)
}

func (suite *fishCommandTest) createRoomRepositoryWithRoom(
	x int,
	y int,
	room *app.Room,
) app.RoomRepository {
	return app.RoomMemoryRepository{}.Create([]*app.Room{room})
}

func (suite *fishCommandTest) createCharacterWithoutTools() *app.Character {
	character := &app.Character{}

	return character
}

func (suite *fishCommandTest) createCharacterWithTool() *app.Character {
	tool := app.Item{}.Create()
	tool.AddFlag(itemFlag.FishTool)
	character := suite.createCharacterWithoutTools()
	character.AddItem(tool)

	return character
}

func (suite *fishCommandTest) createRandomWithSeed(seed int64) *random.Random {
	randomizer := random.Random{}.Create()
	randomizer.Seed(seed)

	return randomizer
}
