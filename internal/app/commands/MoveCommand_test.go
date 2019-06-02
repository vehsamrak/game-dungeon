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

func (suite *moveCommandTest) Test_Execute_CharacterAndDirectionAndGivenCharacterAndRoomRepository_characterMovedIfRoomExists() {
	for id, dataset := range suite.provideCharacterDirectionsAndRooms() {
		roomRepository := dataset.roomRepository
		command := commands.MoveCommand{}.Create(roomRepository)
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
	direction      string
	roomRepository app.RoomRepository
	expectedX      int
	expectedY      int
} {
	getRoomRepositoryWithSingleRoom := func(x int, y int, roomType string) app.RoomRepository {
		return app.RoomMemoryRepository{}.Create([]*app.Room{app.Room{}.Create(x, y, roomType)})
	}

	return []struct {
		direction      string
		roomRepository app.RoomRepository
		expectedX      int
		expectedY      int
	}{
		{"north", getRoomRepositoryWithSingleRoom(0, 0, app.RoomTypeRoad), 0, 0},
		{"south", getRoomRepositoryWithSingleRoom(0, 0, app.RoomTypeRoad), 0, 0},
		{"east", getRoomRepositoryWithSingleRoom(0, 0, app.RoomTypeRoad), 0, 0},
		{"west", getRoomRepositoryWithSingleRoom(0, 0, app.RoomTypeRoad), 0, 0},
		{"north", getRoomRepositoryWithSingleRoom(-1, 0, app.RoomTypeRoad), -1, 0},
		{"south", getRoomRepositoryWithSingleRoom(1, 0, app.RoomTypeRoad), 1, 0},
		{"east", getRoomRepositoryWithSingleRoom(0, 1, app.RoomTypeRoad), 0, 1},
		{"west", getRoomRepositoryWithSingleRoom(0, -1, app.RoomTypeRoad), 0, -1},
		{"north", getRoomRepositoryWithSingleRoom(-1, 0, app.RoomTypeMountain), 0, 0},
	}
}