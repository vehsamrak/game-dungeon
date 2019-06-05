package commands_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
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
	targetRoomX, targetRoomY := 0, 1
	roomBeforeExploration := roomRepository.FindByXY(targetRoomX, targetRoomY)

	err := command.Execute(character, direction.North)

	roomAfterExploration := roomRepository.FindByXY(targetRoomX, targetRoomY)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), roomBeforeExploration)
	assert.NotNil(suite.T(), roomAfterExploration)
}

func (suite *exploreCommandTest) getCharacter() commands.Character {
	return &app.Character{}
}
