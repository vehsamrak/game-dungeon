package app_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"testing"
)

func TestCharacter(test *testing.T) {
	suite.Run(
		test, &characterTest{
			name: "tester",
		},
	)
}

type characterTest struct {
	suite.Suite
	name string
}

func (suite *characterTest) Test_Create_characterParameters_newCharacterCreated() {
	character := app.Character{}.Create(suite.name)

	assert.Equal(suite.T(), suite.name, character.Name())
	assert.Equal(suite.T(), 0, character.X())
	assert.Equal(suite.T(), 0, character.Y())
}

func (suite *characterTest) Test_Move_character_characterMoved() {
	character := suite.createCharacter()
	initialX, initialY := 0, 0
	destinationX, destinationY := 1, 1

	character.Move(destinationX, destinationY)

	assert.NotEqual(suite.T(), initialX, character.X())
	assert.NotEqual(suite.T(), initialY, character.Y())
	assert.Equal(suite.T(), destinationX, character.X())
	assert.Equal(suite.T(), destinationY, character.Y())
}

func (suite *characterTest) Test_HasItemFlag_characterWithItems_validItemDetectionInCharacterInventory() {
	for id, dataset := range suite.provideFlagAndItems() {
		character := suite.createCharacter()
		character.AddItems(dataset.items)

		characterHasItemFlag := character.HasItemFlag(dataset.itemFlag)

		assert.Equal(
			suite.T(),
			dataset.characterHasFlag,
			characterHasItemFlag,
			fmt.Sprintf("Dataset %v %#v", id, dataset),
		)
	}
}

func (suite *characterTest) Test_AddItem_characterWithoutItems_characterMustReveiveItems() {
	character := suite.createCharacter()
	item := app.Item{}.Create()
	items := []*app.Item{item, item}

	character.AddItem(item)
	character.AddItems(items)

	assert.Len(suite.T(), character.Inventory(), 3)
}

func (suite *characterTest) createCharacter() *app.Character {
	return app.Character{}.Create(suite.name)
}

func (suite *characterTest) provideFlagAndItems() []struct {
	itemFlag         string
	items            []*app.Item
	characterHasFlag bool
} {
	axe := app.Item{}.Create()
	axe.AddFlag(app.ItemFlagCutTree)

	tree := app.Item{}.Create()
	tree.AddFlag(app.ItemFlagResourceWood)

	return []struct {
		itemFlag         string
		items            []*app.Item
		characterHasFlag bool
	}{
		{app.ItemFlagCutTree, []*app.Item{axe}, true},
		{app.ItemFlagCutTree, []*app.Item{tree}, false},
		{app.ItemFlagResourceWood, []*app.Item{axe}, false},
		{app.ItemFlagResourceWood, []*app.Item{tree}, true},
		{app.ItemFlagResourceWood, []*app.Item{axe, tree}, true},
	}
}
