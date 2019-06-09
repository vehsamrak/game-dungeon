package commands_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomBiom"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
	"github.com/vehsamrak/game-dungeon/internal/app/random"
	"testing"
)

func TestExploreCommand(test *testing.T) {
	suite.Run(test, &exploreCommandTest{})
}

type exploreCommandTest struct {
	suite.Suite
}

func (suite *exploreCommandTest) Test_Execute_characterAndNoNearRooms_newRoomCreatedWithBiomAndFlagsAndCharacterMovedToNewRoom() {
	for id, dataset := range suite.provideRoomBiomAndFlags() {
		roomRepository := &app.RoomMemoryRepository{}
		command := commands.ExploreCommand{}.Create(roomRepository, suite.createRandomWithSeed(dataset.randomSeed))
		character := suite.createCharacter()
		commandDirection := direction.North
		targetRoomX, targetRoomY := commandDirection.DiffXY()
		roomBeforeExploration := roomRepository.FindByXY(targetRoomX, targetRoomY)

		result := command.Execute(character, commandDirection)

		roomAfterExploration := roomRepository.FindByXY(targetRoomX, targetRoomY)
		assert.False(suite.T(), result.HasErrors(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Nil(suite.T(), roomBeforeExploration, fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.NotNil(suite.T(), roomAfterExploration, fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Equal(suite.T(), dataset.biom, roomAfterExploration.Biom(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		suite.assertRoomHasFlags(dataset.roomFlags, roomAfterExploration, fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Equal(suite.T(), targetRoomX, character.X(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Equal(suite.T(), targetRoomY, character.Y(), fmt.Sprintf("Dataset %v %#v", id, dataset))
	}
}

func (suite *exploreCommandTest) Test_Execute_characterTryToExploreAlreadyExistedRoom_moveCommandExecuted() {
	commandDirection := direction.North
	targetRoomX, targetRoomY := commandDirection.DiffXY()
	room := app.Room{}.Create(targetRoomX, targetRoomY, roomBiom.Forest)
	roomRepository := &app.RoomMemoryRepository{}
	roomRepository.AddRoom(room)
	command := commands.ExploreCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))
	character := suite.createCharacter()
	roomBeforeExploration := roomRepository.FindByXY(targetRoomX, targetRoomY)

	result := command.Execute(character, commandDirection)

	roomAfterExploration := roomRepository.FindByXY(targetRoomX, targetRoomY)
	assert.False(suite.T(), result.HasErrors())
	assert.NotNil(suite.T(), roomBeforeExploration)
	assert.NotNil(suite.T(), roomAfterExploration)
	assert.Equal(suite.T(), targetRoomX, character.X())
	assert.Equal(suite.T(), targetRoomY, character.Y())
}

func (suite *exploreCommandTest) createCharacter() commands.Character {
	return &app.Character{}
}

func (suite *exploreCommandTest) createRandomWithSeed(seed int64) *random.Random {
	randomizer := random.Random{}.Create()
	randomizer.Seed(seed)

	return randomizer
}

func (suite *exploreCommandTest) provideRoomBiomAndFlags() []struct {
	randomSeed int64
	biom       roomBiom.Biom
	roomFlags  []roomFlag.Flag
} {
	return []struct {
		randomSeed int64
		biom       roomBiom.Biom
		roomFlags  []roomFlag.Flag
	}{
		{1, roomBiom.Sea, []roomFlag.Flag{}},
		{2, roomBiom.Mountain, []roomFlag.Flag{roomFlag.OreProbability}},
		{4, roomBiom.Sand, []roomFlag.Flag{}},
		{5, roomBiom.Forest, []roomFlag.Flag{roomFlag.Trees}},
	}
}

func (suite *exploreCommandTest) assertRoomHasFlags(flags []roomFlag.Flag, room *app.Room, message string) {
	for _, flag := range flags {
		assert.True(suite.T(), room.HasFlag(flag), flag.String(), message)
	}
}
