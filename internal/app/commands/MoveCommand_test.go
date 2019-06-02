package commands_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"testing"
)

func TestMoveCommand(test *testing.T) {
	suite.Run(test, &moveCommandTest{})
}

type moveCommandTest struct {
	suite.Suite
}

func (suite *moveCommandTest) Test_Create_noParameters_commandCreated() {
	command := commands.MoveCommand{}.Create()

	assert.NotNil(suite.T(), command)
}

func (suite *moveCommandTest) Test_Execute_CharacterAndDirectionAndGivenCharacterAndGameMap_characterMoved() {
	command := commands.MoveCommand{}.Create()
	character := suite.getCharacter()
	expectedX := character.X() - 1
	expectedY := character.Y()

	command.Execute(character, "north")

	assert.NotNil(suite.T(), command)
	assert.Equal(suite.T(), expectedX, character.X())
	assert.Equal(suite.T(), expectedY, character.Y())
}

func (suite *moveCommandTest) getCharacter() commands.Character {
	return &app.Character{}
}
