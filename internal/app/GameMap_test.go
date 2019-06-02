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
		gameMap := app.GameMap{}.Create(dataset.expectedHeight, dataset.expectedWidth)

		height, width := gameMap.Size()

		assert.NotNil(suite.T(), gameMap)
		assert.Equal(suite.T(), dataset.expectedHeight, height, fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Equal(suite.T(), dataset.expectedWidth, width, fmt.Sprintf("Dataset %v %#v", id, dataset))
	}
}

func (suite *gameMapTest) Test_RoomId_gameMapWithRooms_roomReturned() {
	for id, dataset := range suite.getRooms() {
		gameMap := app.GameMap{}.Create(dataset.x, dataset.y)

		roomId := gameMap.RoomId(dataset.x, dataset.y)

		assert.Equal(suite.T(), dataset.roomId, roomId, fmt.Sprintf("Dataset %v %#v", id, dataset))
	}
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

func (suite *gameMapTest) getRooms() []struct {
	x      int
	y      int
	roomId string
} {
	return []struct {
		x      int
		y      int
		roomId string
	}{
		{1, 1, "1.1"},
		{1, 2, "1.2"},
		{2, 2, "2.2"},
		{12, 23, "12.23"},
		{1, 23, "1.23"},
		{12, 3, "12.3"},
	}
}
