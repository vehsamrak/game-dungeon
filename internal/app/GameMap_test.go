package app_test

import (
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

func (suite *gameMapTest) Test_Create_noParameters_newMapCreated() {
	gameMap := app.GameMap{}.Create()

	assert.NotNil(suite.T(), gameMap)

	// for id, dataset := range suite.dataset() {
	// 	url := appMetricaService.SubstitutePlaceholders(
	// 		dataset.foo,
	// 		dataset.parameters,
	// 	)
	//
	// 	assert.Equal(suite.T(), dataset.bar, url, fmt.Sprintf("dataset #%v", id), dataset)
	// }
}

func (suite *gameMapTest) dataset() []struct {
	foo string
	bar string
} {
	return []struct {
		foo string
		bar string
	}{
		{"", ""},
	}
}
