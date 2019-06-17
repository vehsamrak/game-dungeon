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

	result := command.Execute(character)
	secondResult := command.Execute(character)

	assert.True(suite.T(), result.HasError(gameError.FishNotFound))
	assert.True(suite.T(), secondResult.HasError(gameError.WaitState))
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
	assert.False(suite.T(), character.HasItemFlag(itemFlag.Food))
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
	assert.True(suite.T(), character.HasItemFlag(itemFlag.Food))
}

func (suite *fishCommandTest) createRoomRepositoryWithWaterRoom(
	character *app.Character,
	roomFlags []roomFlag.Flag,
) app.RoomRepository {
	room := app.Room{}.Create(character.X(), character.Y(), character.Z(), roomBiom.Water)
	room.AddFlags(roomFlags)

	return suite.createRoomRepositoryWithRoom(room)
}

func (suite *fishCommandTest) createRoomRepositoryWithForestRoom(character *app.Character) app.RoomRepository {
	room := app.Room{}.Create(character.X(), character.Y(), character.Z(), roomBiom.Forest)

	return suite.createRoomRepositoryWithRoom(room)
}

func (suite *fishCommandTest) createRoomRepositoryWithRoom(room *app.Room) app.RoomRepository {
	return app.RoomMemoryRepository{}.Create([]*app.Room{room})
}

func (suite *fishCommandTest) createCharacterWithoutTools() *app.Character {
	character := app.Character{}.Create("")
	item := character.FindItemWithFlag(itemFlag.FishTool)
	character.DropItem(item)

	return character
}

func (suite *fishCommandTest) createCharacterWithTool() *app.Character {
	tool := app.Item{}.Create()
	tool.AddFlag(itemFlag.FishTool)
	character := suite.createCharacterWithoutTools()
	character.AddItem(tool)

	return character
}

func (suite *fishCommandTest) createRandomWithSeed(seed int64) *random.Randomizer {
	randomizer := random.Randomizer{}.Create()
	randomizer.Seed(seed)

	return randomizer
}
