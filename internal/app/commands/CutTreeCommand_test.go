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
	character := suite.getCharacter()

	noToolError := command.Execute(character)

	assert.NotNil(suite.T(), noToolError)
	assert.Equal(suite.T(), "no tools or room has no trees", noToolError.Error())
}

func (suite *cutTreeCommandTest) getCharacter() commands.Character {
	return &app.Character{}
}
