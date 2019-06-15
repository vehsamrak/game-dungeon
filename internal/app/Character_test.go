package app_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
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
	assert.Equal(suite.T(), 0, character.Z())
}

func (suite *characterTest) Test_Move_character_characterMoved() {
	character := suite.createCharacter()
	initialX, initialY, initialZ := 0, 0, 0
	destinationX, destinationY, destinationZ := 1, 1, 1

	character.Move(destinationX, destinationY, destinationZ)

	assert.NotEqual(suite.T(), initialX, character.X())
	assert.NotEqual(suite.T(), initialY, character.Y())
	assert.NotEqual(suite.T(), initialZ, character.Z())
	assert.Equal(suite.T(), destinationX, character.X())
	assert.Equal(suite.T(), destinationY, character.Y())
	assert.Equal(suite.T(), destinationZ, character.Z())
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

func (suite *characterTest) Test_Health_characterWithHealthPoints_HealthPointsReturned() {
	character := suite.createCharacter()

	healthPoints := character.Health()

	assert.Equal(suite.T(), 100, healthPoints)
}

func (suite *characterTest) createCharacter() *app.Character {
	return app.Character{}.Create(suite.name)
}

func (suite *characterTest) provideFlagAndItems() []struct {
	itemFlag         itemFlag.Flag
	items            []*app.Item
	characterHasFlag bool
} {
	axe := app.Item{}.Create()
	axe.AddFlag(itemFlag.CutTreeTool)

	tree := app.Item{}.Create()
	tree.AddFlag(itemFlag.ResourceWood)

	return []struct {
		itemFlag         itemFlag.Flag
		items            []*app.Item
		characterHasFlag bool
	}{
		{itemFlag.CutTreeTool, []*app.Item{axe}, true},
		{itemFlag.CutTreeTool, []*app.Item{tree}, false},
		{itemFlag.ResourceWood, []*app.Item{axe}, false},
		{itemFlag.ResourceWood, []*app.Item{tree}, true},
		{itemFlag.ResourceWood, []*app.Item{axe, tree}, true},
	}
}
