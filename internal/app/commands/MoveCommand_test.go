package commands_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/exception"
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

		err := command.Execute(character, dataset.direction)

		assert.Equal(suite.T(), dataset.error, err, fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Equal(suite.T(), dataset.expectedCharacterX, character.X(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Equal(suite.T(), dataset.expectedCharacterY, character.Y(), fmt.Sprintf("Dataset %v %#v", id, dataset))
	}
}

func (suite *moveCommandTest) getCharacter() commands.Character {
	return &app.Character{}
}

func (suite *moveCommandTest) provideCharacterDirectionsAndRooms() []struct {
	direction          direction.Direction
	roomRepository     app.RoomRepository
	expectedCharacterX int
	expectedCharacterY int
	error              error
} {
	getRoomRepositoryWithSingleRoom := func(x int, y int, roomFlag string) app.RoomRepository {
		room := app.Room{}.Create(x, y)
		room.AddFlag(roomFlag)

		return app.RoomMemoryRepository{}.Create([]*app.Room{room})
	}

	roomNotFound := exception.RoomNotFound{}
	roomUnfordable := exception.RoomUnfordable{}

	return []struct {
		direction          direction.Direction
		roomRepository     app.RoomRepository
		expectedCharacterX int
		expectedCharacterY int
		error              error
	}{
		{direction.North, getRoomRepositoryWithSingleRoom(0, 0, app.RoomFlagRoad), 0, 0, roomNotFound},
		{direction.South, getRoomRepositoryWithSingleRoom(0, 0, app.RoomFlagRoad), 0, 0, roomNotFound},
		{direction.East, getRoomRepositoryWithSingleRoom(0, 0, app.RoomFlagRoad), 0, 0, roomNotFound},
		{direction.West, getRoomRepositoryWithSingleRoom(0, 0, app.RoomFlagRoad), 0, 0, roomNotFound},
		{direction.North, getRoomRepositoryWithSingleRoom(0, 1, app.RoomFlagRoad), 0, 1, nil},
		{direction.South, getRoomRepositoryWithSingleRoom(0, -1, app.RoomFlagRoad), 0, -1, nil},
		{direction.East, getRoomRepositoryWithSingleRoom(1, 0, app.RoomFlagRoad), 1, 0, nil},
		{direction.West, getRoomRepositoryWithSingleRoom(-1, 0, app.RoomFlagRoad), -1, 0, nil},
		{direction.North, getRoomRepositoryWithSingleRoom(0, 1, app.RoomFlagUnfordable), 0, 0, roomUnfordable},
	}
}
