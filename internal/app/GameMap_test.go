package app_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"testing"
)

func TestGameMap(test *testing.T) {
	suite.Run(test, &gameMapTest{})
}

type gameMapTest struct {
	suite.Suite
}

func (suite *gameMapTest) Test_CreateSize_heightAndWidth_MapCreatedWithExpectedSize() {
	for id, dataset := range suite.getCreateMapSize() {
		gameMap := app.GameMap{}.Create(dataset.expectedHeight, dataset.expectedWidth, suite.getRoomRepository())

		height, width := gameMap.Size()

		assert.NotNil(suite.T(), gameMap)
		assert.Equal(suite.T(), dataset.expectedHeight, height, fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Equal(suite.T(), dataset.expectedWidth, width, fmt.Sprintf("Dataset %v %#v", id, dataset))
	}
}

func (suite *gameMapTest) Test_Room_givenRoomRepositoryWithRoomsAndXYArguments_roomReturned() {
	x, y := 1, 1
	gameMap := app.GameMap{}.Create(x, y, suite.getRoomRepository())

	room := gameMap.Room(x, y)

	assert.Nil(suite.T(), room)
}

func (suite *gameMapTest) getCreateMapSize() []struct {
	expectedHeight int
	expectedWidth  int
} {
	return []struct {
		expectedHeight int
		expectedWidth  int
	}{
		{1, 1},
		{1, 2},
		{2, 2},
		{10, 20},
	}
}

func (suite *gameMapTest) getRoomRepository() app.RoomRepository {
	return &RoomRepositoryMock{}
}

type RoomRepositoryMock struct {
}

func (RoomRepositoryMock) FindByXandY(x int, y int) *app.Room {
	return nil
}

func (RoomRepositoryMock) AddRoom(room *app.Room) {
}

func (RoomRepositoryMock) FindByXY(XY app.XYInterface) *app.Room {
	return nil
}
