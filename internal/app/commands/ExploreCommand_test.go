package commands_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"testing"
)

func TestExploreCommand(test *testing.T) {
	suite.Run(test, &exploreCommandTest{})
}

type exploreCommandTest struct {
	suite.Suite
}

func (suite *exploreCommandTest) Test_Execute_characterAndNoNearRooms_newRoomCreatedWithNewFlags() {
	roomRepository := &app.RoomMemoryRepository{}
	command := commands.ExploreCommand{}.Create(roomRepository)
	character := suite.getCharacter()

	command.Execute(character)
}

func (suite *exploreCommandTest) getCharacter() commands.Character {
	return &app.Character{}
}
