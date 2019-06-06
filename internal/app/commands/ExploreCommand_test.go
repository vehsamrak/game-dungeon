package commands_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/random"
	"testing"
)

func TestExploreCommand(test *testing.T) {
	suite.Run(test, &exploreCommandTest{})
}

type exploreCommandTest struct {
	suite.Suite
}

func (suite *exploreCommandTest) Test_Execute_characterAndNoNearRooms_newRoomCreatedWithNewFlagsCharacterMovedToNewRoom() {
	roomRepository := &app.RoomMemoryRepository{}
	command := commands.ExploreCommand{}.Create(roomRepository, suite.createRandom())
	character := suite.getCharacter()
	targetRoomX, targetRoomY := 0, 1
	roomBeforeExploration := roomRepository.FindByXY(targetRoomX, targetRoomY)

	err := command.Execute(character, direction.North)

	roomAfterExploration := roomRepository.FindByXY(targetRoomX, targetRoomY)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), roomBeforeExploration)
	assert.NotNil(suite.T(), roomAfterExploration)
	assert.True(suite.T(), roomAfterExploration.HasFlag(roomFlag.Trees))
	assert.Equal(suite.T(), targetRoomX, character.X())
	assert.Equal(suite.T(), targetRoomY, character.Y())
}

func (suite *exploreCommandTest) Test_Execute_characterTryToExploreAlreadyExistedRoom_moveCommandExecuted() {
	targetRoomX, targetRoomY := 0, 1
	room := app.Room{}.Create(targetRoomX, targetRoomY)
	roomRepository := &app.RoomMemoryRepository{}
	roomRepository.AddRoom(room)
	command := commands.ExploreCommand{}.Create(roomRepository, suite.createRandom())
	character := suite.getCharacter()
	roomBeforeExploration := roomRepository.FindByXY(targetRoomX, targetRoomY)

	err := command.Execute(character, direction.North)

	roomAfterExploration := roomRepository.FindByXY(targetRoomX, targetRoomY)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), roomBeforeExploration)
	assert.NotNil(suite.T(), roomAfterExploration)
	assert.Equal(suite.T(), targetRoomX, character.X())
	assert.Equal(suite.T(), targetRoomY, character.Y())
}

func (suite *exploreCommandTest) getCharacter() commands.Character {
	return &app.Character{}
}

func (suite *exploreCommandTest) createRandom() *random.Random {
	randomizer := random.Random{}.Create()
	randomizer.Seed(1)

	return randomizer
}
