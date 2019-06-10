package commands_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
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
	allBiomsAreCorrect := true
	for id, dataset := range suite.provideRoomBiomAndFlags() {
		roomRepository := &app.RoomMemoryRepository{}
		command := commands.ExploreCommand{}.Create(roomRepository, suite.createRandomWithSeed(dataset.randomSeed))
		character := suite.createCharacter()
		commandDirection := direction.North
		targetRoomX, targetRoomY := commandDirection.DiffXY()
		roomBeforeExploration := roomRepository.FindByXandY(targetRoomX, targetRoomY)

		result := command.Execute(character, commandDirection)

		roomAfterExploration := roomRepository.FindByXandY(targetRoomX, targetRoomY)
		assert.False(suite.T(), result.HasErrors(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Nil(suite.T(), roomBeforeExploration, fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.NotNil(suite.T(), roomAfterExploration, fmt.Sprintf("Dataset %v %#v", id, dataset))
		biomIsCorrect := assert.Equal(suite.T(), dataset.biom, roomAfterExploration.Biom(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		if !biomIsCorrect {
			allBiomsAreCorrect = false
		}
		suite.assertRoomHasFlags(dataset.roomFlags, roomAfterExploration, fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Equal(suite.T(), targetRoomX, character.X(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Equal(suite.T(), targetRoomY, character.Y(), fmt.Sprintf("Dataset %v %#v", id, dataset))
	}

	if !allBiomsAreCorrect {
		suite.showBiomNumbers(int64(len(roomBiom.All()) * 10))
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
	roomBeforeExploration := roomRepository.FindByXandY(targetRoomX, targetRoomY)

	result := command.Execute(character, commandDirection)

	roomAfterExploration := roomRepository.FindByXandY(targetRoomX, targetRoomY)
	assert.False(suite.T(), result.HasErrors())
	assert.NotNil(suite.T(), roomBeforeExploration)
	assert.NotNil(suite.T(), roomAfterExploration)
	assert.Equal(suite.T(), targetRoomX, character.X())
	assert.Equal(suite.T(), targetRoomY, character.Y())
}

func (suite *exploreCommandTest) Test_Execute_characterInCaveBiomAndNoNearRooms_wrongBiom() {
	character := suite.createCharacter()
	room := app.Room{}.Create(character.X(), character.Y(), roomBiom.Cave)
	roomRepository := &app.RoomMemoryRepository{}
	roomRepository.AddRoom(room)
	commandDirection := direction.North
	targetRoomX, targetRoomY := commandDirection.DiffXY()
	command := commands.ExploreCommand{}.Create(roomRepository, suite.createRandomWithSeed(0))

	result := command.Execute(character, commandDirection)

	assert.True(suite.T(), result.HasError(gameError.WrongBiom))
	assert.Nil(suite.T(), roomRepository.FindByXandY(targetRoomX, targetRoomY))
	assert.NotEqual(suite.T(), targetRoomX+targetRoomY, character.X()+character.Y())
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
		{18, roomBiom.Swamp, []roomFlag.Flag{}},
		{0, roomBiom.Hill, []roomFlag.Flag{roomFlag.CaveProbability}},
		{13, roomBiom.Water, []roomFlag.Flag{roomFlag.FishProbability}},
		{14, roomBiom.Cave, []roomFlag.Flag{roomFlag.OreProbability, roomFlag.GemProbability}},
		{7, roomBiom.Plain, []roomFlag.Flag{}},
		{3, roomBiom.Cliff, []roomFlag.Flag{roomFlag.Unfordable}},
		{1, roomBiom.Sand, []roomFlag.Flag{roomFlag.GemProbability}},
		{2, roomBiom.Town, []roomFlag.Flag{}},
		{5, roomBiom.Air, []roomFlag.Flag{}},
		{11, roomBiom.Forest, []roomFlag.Flag{roomFlag.Trees}},
		{26, roomBiom.Mountain, []roomFlag.Flag{roomFlag.CaveProbability}},
	}
}

func (suite *exploreCommandTest) showBiomNumbers(iterationsCount int64) {
	var i int64
	for i = 0; i < iterationsCount; i++ {
		character := suite.createCharacter()
		roomRepository := &app.RoomMemoryRepository{}
		command := commands.ExploreCommand{}.Create(roomRepository, suite.createRandomWithSeed(i))
		commandDirection := direction.North
		command.Execute(character, commandDirection)
		targetRoomX, targetRoomY := commandDirection.DiffXY()
		roomAfterExploration := roomRepository.FindByXandY(targetRoomX, targetRoomY)

		fmt.Printf("%v %v\n", i, roomAfterExploration.Biom())
	}
}

func (suite *exploreCommandTest) assertRoomHasFlags(flags []roomFlag.Flag, room *app.Room, message string) {
	for _, flag := range flags {
		assert.True(suite.T(), room.HasFlag(flag), flag.String(), message)
	}
}
