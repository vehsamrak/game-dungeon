package commands_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"testing"
)

func TestMoveCommand(test *testing.T) {
	suite.Run(test, &moveCommandTest{})
}

type moveCommandTest struct {
	suite.Suite
}

func (suite *moveCommandTest) Test_Execute_CharacterAndDirectionAndGivenCharacterAndRoomRepository_characterMovedIfRoomExists() {
	for id, dataset := range suite.provideCharacterDirectionsAndRooms() {
		command := commands.MoveCommand{}.Create(dataset.roomRepository)
		character := suite.getCharacter()

		command.Execute(character, dataset.direction)

		assert.Equal(suite.T(), dataset.expectedX, character.X(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Equal(suite.T(), dataset.expectedY, character.Y(), fmt.Sprintf("Dataset %v %#v", id, dataset))
	}
}

func (suite *moveCommandTest) getCharacter() commands.Character {
	return &app.Character{}
}

func (suite *moveCommandTest) provideCharacterDirectionsAndRooms() []struct {
	direction      direction.Direction
	roomRepository app.RoomRepository
	expectedX      int
	expectedY      int
} {
	getRoomRepositoryWithSingleRoom := func(x int, y int, roomFlag string) app.RoomRepository {
		room := app.Room{}.Create(x, y)
		room.AddFlag(roomFlag)

		return app.RoomMemoryRepository{}.Create([]*app.Room{room})
	}

	return []struct {
		direction      direction.Direction
		roomRepository app.RoomRepository
		expectedX      int
		expectedY      int
	}{
		{direction.North, getRoomRepositoryWithSingleRoom(0, 0, app.RoomFlagRoad), 0, 0},
		{direction.South, getRoomRepositoryWithSingleRoom(0, 0, app.RoomFlagRoad), 0, 0},
		{direction.East, getRoomRepositoryWithSingleRoom(0, 0, app.RoomFlagRoad), 0, 0},
		{direction.West, getRoomRepositoryWithSingleRoom(0, 0, app.RoomFlagRoad), 0, 0},
		{direction.North, getRoomRepositoryWithSingleRoom(-1, 0, app.RoomFlagRoad), -1, 0},
		{direction.South, getRoomRepositoryWithSingleRoom(1, 0, app.RoomFlagRoad), 1, 0},
		{direction.East, getRoomRepositoryWithSingleRoom(0, 1, app.RoomFlagRoad), 0, 1},
		{direction.West, getRoomRepositoryWithSingleRoom(0, -1, app.RoomFlagRoad), 0, -1},
		{direction.North, getRoomRepositoryWithSingleRoom(-1, 0, app.RoomFlagUnfordable), 0, 0},
	}
}
