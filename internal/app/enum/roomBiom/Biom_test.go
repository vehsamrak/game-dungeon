package roomBiom_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomBiom"
	"testing"
)

func TestBiom(test *testing.T) {
	suite.Run(test, &biomTest{})
}

type biomTest struct {
	suite.Suite
}

func (suite *biomTest) Test_ExploreProbability_biom_exploreProbabilityPercentage() {
	for id, dataset := range suite.provideBiomsAndExploreProbabilities() {

		exploreProbability := dataset.biom.ExploreProbability()

		assert.Equal(suite.T(), dataset.probability, exploreProbability, fmt.Sprintf("Dataset %v %#v", id, dataset))
	}
}

func (suite *biomTest) provideBiomsAndExploreProbabilities() []struct {
	biom        roomBiom.Biom
	probability int
} {
	return []struct {
		biom        roomBiom.Biom
		probability int
	}{
		{roomBiom.Forest, 15},
		{roomBiom.Plain, 15},
		{roomBiom.Hill, 15},
		{roomBiom.Water, 10},
		{roomBiom.Swamp, 7},
		{roomBiom.Sand, 5},
		{roomBiom.Mountain, 5},
		{roomBiom.Cave, 2},
		{roomBiom.Cliff, 1},
		{roomBiom.Town, 1},
	}
}
