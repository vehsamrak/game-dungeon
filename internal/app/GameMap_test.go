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

func (suite *gameMapTest) Test_Create_heightAndWidth_MapCreatedWithHeightAndWidth() {
	for id, dataset := range suite.dataset() {
		gameMap := app.GameMap{}.Create(dataset.expectedHeight, dataset.expectedWidth)
		height, width := gameMap.GetHeightAndWidth()

		assert.NotNil(suite.T(), gameMap)
		assert.Equal(suite.T(), dataset.expectedHeight, height, fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Equal(suite.T(), dataset.expectedWidth, width, fmt.Sprintf("Dataset %v %#v", id, dataset))
	}
}

func (suite *gameMapTest) dataset() []struct {
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
	}
}
