package commands_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomBiom"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
	"testing"
)

func TestMoveCommand(test *testing.T) {
	suite.Run(test, &moveCommandTest{})
}

type moveCommandTest struct {
	suite.Suite
}

func (suite *moveCommandTest) Test_Execute_CharacterAndDirectionAndRoomRepository_characterMovedIfRoomExists() {
	for id, dataset := range suite.provideCharacterDirectionsAndRooms() {
		command := commands.MoveCommand{}.Create(dataset.roomRepository)
		character := suite.getCharacter()

		result := command.Execute(character, dataset.direction.String())

		if dataset.error != "" {
			assert.True(suite.T(), result.HasError(dataset.error), fmt.Sprintf("Dataset %v %#v", id, dataset))
		}
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
	expectedCharacterZ int
	error              gameError.Error
} {
	getRoomRepositoryWithSingleRoom := func(x int, y int, z int, roomFlag roomFlag.Flag) app.RoomRepository {
		room := app.Room{}.Create(x, y, z, roomBiom.Forest)
		room.AddFlag(roomFlag)

		return app.RoomMemoryRepository{}.Create([]*app.Room{room})
	}

	roomNotFound := gameError.RoomNotFound
	roomUnfordable := gameError.RoomUnfordable

	return []struct {
		direction          direction.Direction
		roomRepository     app.RoomRepository
		expectedCharacterX int
		expectedCharacterY int
		expectedCharacterZ int
		error              gameError.Error
	}{
		{direction.North, getRoomRepositoryWithSingleRoom(0, 0, 0, roomFlag.Road), 0, 0, 0, roomNotFound},
		{direction.South, getRoomRepositoryWithSingleRoom(0, 0, 0, roomFlag.Road), 0, 0, 0, roomNotFound},
		{direction.East, getRoomRepositoryWithSingleRoom(0, 0, 0, roomFlag.Road), 0, 0, 0, roomNotFound},
		{direction.West, getRoomRepositoryWithSingleRoom(0, 0, 0, roomFlag.Road), 0, 0, 0, roomNotFound},
		{direction.North, getRoomRepositoryWithSingleRoom(0, 1, 0, roomFlag.Road), 0, 1, 0, ""},
		{direction.South, getRoomRepositoryWithSingleRoom(0, -1, 0, roomFlag.Road), 0, -1, 0, ""},
		{direction.East, getRoomRepositoryWithSingleRoom(1, 0, 0, roomFlag.Road), 1, 0, 0, ""},
		{direction.West, getRoomRepositoryWithSingleRoom(-1, 0, 0, roomFlag.Road), -1, 0, 0, ""},
		{direction.Up, getRoomRepositoryWithSingleRoom(0, 0, 1, roomFlag.Road), 0, 0, 1, ""},
		{direction.Down, getRoomRepositoryWithSingleRoom(0, 0, -1, roomFlag.Road), 0, 0, -1, ""},
		{direction.North, getRoomRepositoryWithSingleRoom(0, 1, 0, roomFlag.Unfordable), 0, 0, 0, roomUnfordable},
	}
}
