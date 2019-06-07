package commands_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/exception"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
	"testing"
)

func TestCutTreeCommand(test *testing.T) {
	suite.Run(test, &cutTreeCommandTest{})
}

type cutTreeCommandTest struct {
	suite.Suite
}

func (suite *cutTreeCommandTest) Test_Execute_characterWithoutTool_noToolError() {
	roomRepository := &app.RoomMemoryRepository{}
	command := commands.CutTreeCommand{}.Create(roomRepository)
	character := suite.getCharacter(nil)

	noToolError := command.Execute(character)

	assert.NotNil(suite.T(), noToolError)
	assert.Equal(suite.T(), exception.NoTool{}, noToolError)
	assert.NotEmpty(suite.T(), noToolError.Error())
}

func (suite *cutTreeCommandTest) Test_Execute_characterWithToolAndRoomHasTrees_treeAppearsInCharacterInventory() {
	axe := app.Item{}.Create()
	axe.AddFlag(itemFlag.CutTree)
	character := &app.Character{}
	character.AddItem(axe)
	roomRepository := suite.getRoomRepositoryWithSingleRoom(
		character.X(),
		character.Y(),
		[]roomFlag.Flag{roomFlag.Trees},
	)
	command := commands.CutTreeCommand{}.Create(roomRepository)
	characterItemsCountBeforeCommand := len(character.Inventory())
	characterHasWoodBeforeCommand := character.HasItemFlag(itemFlag.ResourceWood)

	err := command.Execute(character)

	characterItemsCountAfterCommand := len(character.Inventory())
	characterHasWoodAfterCommand := character.HasItemFlag(itemFlag.ResourceWood)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, characterItemsCountBeforeCommand)
	assert.Equal(suite.T(), 2, characterItemsCountAfterCommand)
	assert.False(suite.T(), characterHasWoodBeforeCommand)
	assert.True(suite.T(), characterHasWoodAfterCommand)
}

func (suite *cutTreeCommandTest) getCharacter(items []*app.Item) commands.Character {
	character := &app.Character{}
	character.AddItems(items)

	return character
}

func (suite *cutTreeCommandTest) getRoomRepositoryWithSingleRoom(x int, y int, roomFlags []roomFlag.Flag) app.RoomRepository {
	room := app.Room{}.Create(x, y)
	room.AddFlags(roomFlags)

	return app.RoomMemoryRepository{}.Create([]*app.Room{room})
}
