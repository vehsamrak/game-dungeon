package commands_test

import (
	"fmt"
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

func (suite *moveCommandTest) Test_Execute_CharacterAndDirectionAndGivenCharacter_characterMoved() {
	command := commands.MoveCommand{}.Create()

	for id, dataset := range suite.provideCharacterDirections() {
		character := suite.getCharacter()

		command.Execute(character, dataset.direction)

		assert.Equal(suite.T(), dataset.expectedX, character.X(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Equal(suite.T(), dataset.expectedY, character.Y(), fmt.Sprintf("Dataset %v %#v", id, dataset))
	}
}

func (suite *moveCommandTest) getCharacter() commands.Character {
	return &app.Character{}
}

func (suite *moveCommandTest) provideCharacterDirections() []struct {
	direction string
	expectedX int
	expectedY int
} {
	return []struct {
		direction string
		expectedX int
		expectedY int
	}{
		{"north", -1, 0},
		{"south", 1, 0},
		{"east", 0, 1},
		{"west", 0, -1},
	}
}
