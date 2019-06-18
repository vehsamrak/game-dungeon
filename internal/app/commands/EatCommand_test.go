package commands_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
	"testing"
)

func TestEatCommand(test *testing.T) {
	suite.Run(test, &eatCommandTest{})
}

type eatCommandTest struct {
	suite.Suite
}

func (suite *eatCommandTest) Test_HealthPrice_noParameters_zeroPriceReturned() {
	command := commands.EatCommand{}.Create()

	healthPrice := command.HealthPrice()

	assert.Equal(suite.T(), 0, healthPrice)
}

func (suite *eatCommandTest) Test_Execute_characterWithoutFood_foodNotFoundError() {
	command := commands.EatCommand{}.Create()
	character := suite.createCharacter(&app.Item{})

	result := command.Execute(character)

	assert.True(suite.T(), result.HasError(gameError.FoodNotFound))
}

func (suite *eatCommandTest) Test_Execute_characterInventory_foodRemovedFromInventoryAndHPIncreased() {
	for id, dataset := range suite.provideFoodAndCharacterHP() {
		command := commands.EatCommand{}.Create()
		character := suite.createCharacter(dataset.food)
		character.LowerHealth(character.MaxHealth())
		character.IncreaseHealth(dataset.initialHP)

		result := command.Execute(character)

		assert.False(suite.T(), result.HasErrors(), fmt.Sprintf("Dataset %v %#v", id, dataset))
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

func (suite *eatCommandTest) createCharacter(item *app.Item) *app.Character {
	character := app.Character{}.Create("")
	character.AddItem(item)

	return character
}

func (suite *eatCommandTest) provideFoodAndCharacterHP() []struct {
	food        *app.Item
	initialHP   int
	increasedHP int
} {
	food := app.Item{}.Create("")
	food.AddFlag(itemFlag.Food)

	return []struct {
		food        *app.Item
		initialHP   int
		increasedHP int
	}{
		{food, 0, 10},
		{food, 89, 99},
		{food, 90, 100},
		{food, 91, 100},
		{food, 95, 100},
		{food, 99, 100},
	}
}
