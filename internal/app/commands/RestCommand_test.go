package commands_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/timer"
	"testing"
	"time"
)

func TestRestCommand(test *testing.T) {
	suite.Run(test, &restCommandTest{})
}

type restCommandTest struct {
	suite.Suite
}

func (suite *restCommandTest) Test_HealthPrice_noParameters_zeroPriceReturned() {
	command := commands.RestCommand{}.Create()

	healthPrice := command.HealthPrice()

	assert.Equal(suite.T(), 0, healthPrice)
}

func (suite *restCommandTest) Test_Execute_character_WaitStateErrorOrHPIncreased() {
	command := commands.RestCommand{}.Create()
	for id, dataset := range suite.provideLastRestTimeAndInitialHPAndIncreasedHP() {
		character := suite.createCharacterWithHPAndTimer(dataset.initialHP, dataset.lastRest)

		result := command.Execute(character)

		assert.Equal(suite.T(), dataset.error != "", result.HasErrors(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		if dataset.error != "" {
			assert.True(suite.T(), result.HasError(dataset.error))
		}
		assert.Equal(
			suite.T(),
			dataset.increasedHP,
			character.Health(),
			fmt.Sprintf("Dataset %v %#v", id, dataset),
		)
		assert.True(
			suite.T(),
			character.Health() <= character.MaxHealth(),
			fmt.Sprintf("Dataset %v %#v", id, dataset),
		)
	}
}

func (suite *restCommandTest) createCharacterWithHPAndTimer(hp int, lastRest time.Duration) *app.Character {
	character := app.Character{}.Create("")
	character.LowerHealth(character.MaxHealth())
	character.IncreaseHealth(hp)
	character.SetTimer(timer.Rest, lastRest)

	return character
}

func (suite *restCommandTest) provideLastRestTimeAndInitialHPAndIncreasedHP() []struct {
	lastRest    time.Duration
	initialHP   int
	increasedHP int
	error       gameError.Error
} {
	return []struct {
		lastRest    time.Duration
		initialHP   int
		increasedHP int
		error       gameError.Error
	}{
		{1 * time.Second, 90, 90, gameError.WaitState},
		{0 * time.Second, 90, 92, ""},
		{-1 * time.Second, 90, 92, ""},
	}
}
