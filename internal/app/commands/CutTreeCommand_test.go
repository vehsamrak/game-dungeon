package commands_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
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
	assert.Equal(suite.T(), "no tools or room has no trees", noToolError.Error())
}

func (suite *cutTreeCommandTest) Test_Execute_characterWithToolAndRoomHasTrees_treeAppearsInCharacterInventory() {
	axe := app.Item{}.Create()
	axe.AddFlag(app.ItemFlagCutTree)
	character := &app.Character{}
	character.AddItem(axe)
	roomRepository := suite.getRoomRepositoryWithSingleRoom(character.X(), character.Y(), []string{app.RoomFlagHasTrees})
	command := commands.CutTreeCommand{}.Create(roomRepository)
	characterItemsCountBeforeCommand := len(character.Inventory())
	characterHasWoodBeforeCommand := character.HasItemFlag(app.ItemFlagResourceWood)

	err := command.Execute(character)

	characterItemsCountAfterCommand := len(character.Inventory())
	characterHasWoodAfterCommand := character.HasItemFlag(app.ItemFlagResourceWood)
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

func (suite *cutTreeCommandTest) getRoomRepositoryWithSingleRoom(x int, y int, roomFlags []string) app.RoomRepository {
	room := app.Room{}.Create(x, y)
	room.AddFlags(roomFlags)

	return app.RoomMemoryRepository{}.Create([]*app.Room{room})
}
