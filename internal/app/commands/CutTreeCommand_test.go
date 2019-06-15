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
	character := suite.createCharacter(nil)

	result := command.Execute(character)

	assert.True(suite.T(), result.HasError(gameError.NoTool))
}

func (suite *cutTreeCommandTest) Test_Execute_characterWithToolAndRoomHasTrees_treeAppearsInCharacterInventory() {
	axe := app.Item{}.Create()
	axe.AddFlag(itemFlag.CutTreeTool)
	character := &app.Character{}
	character.AddItem(axe)
	roomRepository := suite.createRoomRepositoryWithRoom(
		character.X(),
		character.Y(),
		character.Z(),
		[]roomFlag.Flag{roomFlag.Trees},
	)
	command := commands.CutTreeCommand{}.Create(roomRepository)
	characterInventoryBeforeCommand := character.Inventory()
	characterHasWoodBeforeCommand := character.HasItemFlag(itemFlag.ResourceWood)

	command.Execute(character)

	characterHasWoodAfterCommand := character.HasItemFlag(itemFlag.ResourceWood)
	assert.Len(suite.T(), characterInventoryBeforeCommand, 1)
	assert.Len(suite.T(), character.Inventory(), 2)
	assert.False(suite.T(), characterHasWoodBeforeCommand)
	assert.True(suite.T(), characterHasWoodAfterCommand)
}

func (suite *cutTreeCommandTest) createCharacter(items []*app.Item) commands.Character {
	character := &app.Character{}
	character.AddItems(items)

	return character
}

func (suite *cutTreeCommandTest) createRoomRepositoryWithRoom(
	x int,
	y int,
	z int,
	roomFlags []roomFlag.Flag,
) app.RoomRepository {
	room := app.Room{}.Create(x, y, z, roomBiom.Forest)
	room.AddFlags(roomFlags)

	return app.RoomMemoryRepository{}.Create([]*app.Room{room})
}
