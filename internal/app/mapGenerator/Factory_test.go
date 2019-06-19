package mapGenerator_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app/mapGenerator"
	"testing"
)

func TestFactory(test *testing.T) {
	suite.Run(test, &factoryTest{})
}

type factoryTest struct {
	suite.Suite
}

func (suite *factoryTest) Test_CreateGenerator_generatorName_specificMapGeneratorCreated() {
	for id, dataset := range []struct {
		generatorName mapGenerator.Name
	}{
		{mapGenerator.Water},
	} {
		factory := mapGenerator.Factory{}.Create()

		generator := factory.CreateGenerator(dataset.generatorName)

		assert.NotNil(suite.T(), generator, fmt.Sprintf("Dataset %v %#v", id, dataset))
	}
}
