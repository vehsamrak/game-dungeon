package commands_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/exception"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/notice"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/random"
	"testing"
)

func TestSearchOreCommand(test *testing.T) {
	suite.Run(test, &searchOreCommandTest{})
}

type searchOreCommandTest struct {
	suite.Suite
}

func (suite *searchOreCommandTest) Test_Execute_characterWithoutTool_noTool() {
	character := suite.createCharacterWithoutTools()
	roomRepository := &app.RoomMemoryRepository{}
	command := commands.SearchOreCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))

	commandResult := command.Execute(character)

	assert.True(suite.T(), commandResult.HasError(exception.NoTool{}))
}

func (suite *searchOreCommandTest) Test_Execute_characterWithToolAndRoomHasNoOreProbability_oreNotFound() {
	character := suite.createCharacterWithTool()
	roomRepository := suite.createRoomRepositoryWithRoom(
		character.X(),
		character.Y(),
		[]roomFlag.Flag{},
	)
	command := commands.SearchOreCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))

	commandResult := command.Execute(character)

	assert.True(suite.T(), commandResult.HasError(exception.OreNotFound{}))
}

func (suite *searchOreCommandTest) Test_Execute_characterWithToolAndRoomHasOreProbabilityButNoOre_oreNotFound() {
	character := suite.createCharacterWithTool()
	roomRepository := suite.createRoomRepositoryWithRoom(
		character.X(),
		character.Y(),
		[]roomFlag.Flag{roomFlag.OreProbability},
	)
	command := commands.SearchOreCommand{}.Create(roomRepository, suite.createRandomWithSeed(0))

	commandResult := command.Execute(character)

	assert.True(suite.T(), commandResult.HasError(exception.OreNotFound{}))
}

func (suite *searchOreCommandTest) Test_Execute_characterWithToolAndRoomHasOreProbabilityButNoOre_oreFound() {
	character := suite.createCharacterWithTool()
	characterBeforeCommand := *character
	roomRepository := suite.createRoomRepositoryWithRoom(
		character.X(),
		character.Y(),
		[]roomFlag.Flag{roomFlag.OreProbability},
	)
	command := commands.SearchOreCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))

	commandResult := command.Execute(character)

	assert.False(suite.T(), commandResult.HasErrors())
	assert.Equal(suite.T(), character.Inventory(), characterBeforeCommand.Inventory())
	assert.False(suite.T(), character.HasItemFlag(itemFlag.ResourceOre))
	assert.True(suite.T(), commandResult.HasNotice(notice.FoundOre))
}

func (suite *searchOreCommandTest) createRoomRepositoryWithRoom(x int, y int, roomFlags []roomFlag.Flag) app.RoomRepository {
	room := app.Room{}.Create(x, y)
	room.AddFlags(roomFlags)

	return app.RoomMemoryRepository{}.Create([]*app.Room{room})
}

func (suite *searchOreCommandTest) createCharacterWithoutTools() *app.Character {
	character := &app.Character{}

	return character
}

func (suite *searchOreCommandTest) createCharacterWithTool() *app.Character {
	tool := app.Item{}.Create()
	tool.AddFlag(itemFlag.SearchOre)
	character := &app.Character{}
	character.AddItem(tool)

	return character
}

func (suite *searchOreCommandTest) createRandomWithSeed(seed int64) *random.Random {
	randomizer := random.Random{}.Create()
	randomizer.Seed(seed)

	return randomizer
}
