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

func (suite *exploreCommandTest) Test_Execute_characterAndNoNearRooms_newRoomsCreatedWithNewFlagsCharacterMovedToNewRoom() {
	roomRepository := &app.RoomMemoryRepository{}
	command := commands.ExploreCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))
	character := suite.getCharacter()
	firstExploredRoomX, firstExploredRoomY := 0, 1
	secondExploredRoomX, secondExploredRoomY := 0, 2
	firstRoomBeforeExploration := roomRepository.FindByXY(firstExploredRoomX, firstExploredRoomY)
	secondRoomBeforeExploration := roomRepository.FindByXY(secondExploredRoomX, secondExploredRoomY)

	firstRoomError := command.Execute(character, direction.North)
	firstRoomAfterExploration := roomRepository.FindByXY(firstExploredRoomX, firstExploredRoomY)
	secondRoomError := command.Execute(character, direction.North)
	secondRoomAfterExploration := roomRepository.FindByXY(secondExploredRoomX, secondExploredRoomY)

	assert.Nil(suite.T(), firstRoomError)
	assert.Nil(suite.T(), secondRoomError)
	assert.Nil(suite.T(), firstRoomBeforeExploration)
	assert.Nil(suite.T(), secondRoomBeforeExploration)
	assert.NotNil(suite.T(), firstRoomAfterExploration)
	assert.NotNil(suite.T(), secondRoomAfterExploration)
	assert.True(suite.T(), firstRoomAfterExploration.HasFlag(roomFlag.Unfordable))
	assert.True(suite.T(), secondRoomAfterExploration.HasFlag(roomFlag.OreProbability))
	assert.Equal(suite.T(), secondExploredRoomX, character.X())
	assert.Equal(suite.T(), secondExploredRoomY, character.Y())
}

func (suite *exploreCommandTest) Test_Execute_characterTryToExploreAlreadyExistedRoom_moveCommandExecuted() {
	targetRoomX, targetRoomY := 0, 1
	room := app.Room{}.Create(targetRoomX, targetRoomY)
	roomRepository := &app.RoomMemoryRepository{}
	roomRepository.AddRoom(room)
	command := commands.ExploreCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))
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

func (suite *exploreCommandTest) createRandomWithSeed(seed int64) *random.Random {
	randomizer := random.Random{}.Create()
	randomizer.Seed(seed)

	return randomizer
}
