package commands_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/commands"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
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

func (suite *exploreCommandTest) Test_Execute_characterCanFlyAndNoNearRooms_newRoomCreatedWithBiomAndCharacterMovedToNewRoom() {
	allBiomsAreCorrect := true
	for id, dataset := range suite.provideRoomBioms() {
		character := suite.createCharacter()
		character.Move(0, 0, dataset.z)
		flyItem := app.Item{}.Create()
		flyItem.AddFlag(itemFlag.CanFly)
		character.AddItem(flyItem)
		roomRepository := &app.RoomMemoryRepository{}
		roomRepository.AddRoom(app.Room{}.Create(character.X(), character.Y(), character.Z(), roomBiom.Forest))
		command := commands.ExploreCommand{}.Create(roomRepository, suite.createRandomWithSeed(dataset.randomSeed))
		commandDirection := direction.North
		directionRoomX, directionRoomY, directionRoomZ := commandDirection.DiffXYZ()
		targetRoomX := directionRoomX + character.X()
		targetRoomY := directionRoomY + character.Y()
		targetRoomZ := directionRoomZ + character.Z()
		roomRepository.AddRoom(app.Room{}.Create(targetRoomX, targetRoomY, targetRoomZ-1, roomBiom.Mountain))
		roomBeforeExploration := roomRepository.FindByXYandZ(targetRoomX, targetRoomY, targetRoomZ)

		result := command.Execute(character, commandDirection.String())

		roomAfterExploration := roomRepository.FindByXYandZ(targetRoomX, targetRoomY, targetRoomZ)
		assert.False(suite.T(), result.HasErrors(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Nil(suite.T(), roomBeforeExploration, fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.NotNil(suite.T(), roomAfterExploration, fmt.Sprintf("Dataset %v %#v", id, dataset))
		biomIsCorrect := assert.Equal(
			suite.T(),
			dataset.biom,
			roomAfterExploration.Biom(),
			fmt.Sprintf("Dataset %v %#v", id, dataset),
		)
		if !biomIsCorrect {
			allBiomsAreCorrect = false
		}
		assert.Equal(suite.T(), targetRoomX, character.X(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Equal(suite.T(), targetRoomY, character.Y(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Equal(suite.T(), targetRoomZ, character.Z(), fmt.Sprintf("Dataset %v %#v", id, dataset))
	}
	if !allBiomsAreCorrect {
		suite.showBiomNumbers(0, len(roomBiom.All())-1) // all bioms except air
		suite.showBiomNumbers(1, len([]roomBiom.Biom{roomBiom.Air, roomBiom.Mountain, roomBiom.Cliff}))
	}
}

func (suite *exploreCommandTest) Test_Execute_characterTryToExploreAlreadyExistedRoom_moveCommandExecuted() {
	commandDirection := direction.North
	character := suite.createCharacter()
	targetRoomX, targetRoomY, targetRoomZ := commandDirection.DiffXYZ()
	initialRoom := app.Room{}.Create(character.X(), character.Y(), character.Z(), roomBiom.Forest)
	destinationRoom := app.Room{}.Create(targetRoomX, targetRoomY, targetRoomZ, roomBiom.Forest)
	roomRepository := app.RoomMemoryRepository{}.Create([]*app.Room{initialRoom, destinationRoom})
	command := commands.ExploreCommand{}.Create(roomRepository, suite.createRandomWithSeed(1))
	roomBeforeExploration := roomRepository.FindByXYandZ(targetRoomX, targetRoomY, targetRoomZ)

	result := command.Execute(character, commandDirection.String())
	secondResult := command.Execute(character, commandDirection.String())

	roomAfterExploration := roomRepository.FindByXYandZ(targetRoomX, targetRoomY, targetRoomZ)
	assert.False(suite.T(), result.HasErrors())
	assert.NotNil(suite.T(), roomBeforeExploration)
	assert.NotNil(suite.T(), roomAfterExploration)
	assert.Equal(suite.T(), targetRoomX, character.X())
	assert.Equal(suite.T(), targetRoomY, character.Y())
	assert.True(suite.T(), secondResult.HasError(gameError.WaitState))
}

func (suite *exploreCommandTest) Test_Execute_characterInDisallowedFromExploreBiomAndNoNearRooms_wrongBiom() {
	for id, dataset := range suite.provideDisallowedFromExploreBioms() {
		character := suite.createCharacter()
		room := app.Room{}.Create(character.X(), character.Y(), character.Z(), dataset.biom)
		roomRepository := app.RoomMemoryRepository{}.Create([]*app.Room{room})
		commandDirection := direction.North
		targetRoomX, targetRoomY, targetRoomZ := commandDirection.DiffXYZ()
		command := commands.ExploreCommand{}.Create(roomRepository, suite.createRandomWithSeed(0))

		result := command.Execute(character, commandDirection.String())

		assert.True(suite.T(), result.HasError(gameError.WrongBiom), fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Nil(
			suite.T(),
			roomRepository.FindByXYandZ(
				targetRoomX,
				targetRoomY,
				targetRoomZ,
			),
			fmt.Sprintf("Dataset %v %#v", id, dataset),
		)
		assert.NotEqual(
			suite.T(),
			targetRoomX+targetRoomY, character.X()+character.Y(), fmt.Sprintf("Dataset %v %#v", id, dataset),
		)
	}
}

func (suite *exploreCommandTest) Test_Execute_characterInExplorableBiomAndNoNearRoomsAndWrongDirection_wrongDirection() {
	for id, dataset := range suite.provideDisallowedDirections() {
		character := suite.createCharacter()
		room := app.Room{}.Create(character.X(), character.Y(), character.Z(), roomBiom.Forest)
		roomRepository := app.RoomMemoryRepository{}.Create([]*app.Room{room})
		command := commands.ExploreCommand{}.Create(roomRepository, suite.createRandomWithSeed(0))

		result := command.Execute(character, dataset.direction.String())

		assert.True(suite.T(), result.HasError(gameError.WrongDirection), fmt.Sprintf("Dataset %v %#v", id, dataset))
	}
}

func (suite *exploreCommandTest) Test_Execute_characterCanExploreAir_airExploredAndCharacterMovedBackIfCantFly() {
	allBiomsAreCorrect := true
	for id, dataset := range suite.provideFlyItem() {
		character := suite.createCharacter()
		character.Move(0, 0, 1)
		flyItem := app.Item{}.Create()
		flyItem.AddFlag(dataset.itemFlag)
		character.AddItem(flyItem)
		roomRepository := &app.RoomMemoryRepository{}
		initialRoom := app.Room{}.Create(character.X(), character.Y(), character.Z(), roomBiom.Mountain)
		roomRepository.AddRoom(initialRoom)
		command := commands.ExploreCommand{}.Create(roomRepository, suite.createRandomWithSeed(18))
		commandDirection := direction.North
		directionRoomX, directionRoomY, directionRoomZ := commandDirection.DiffXYZ()
		targetRoomX := directionRoomX + character.X()
		targetRoomY := directionRoomY + character.Y()
		targetRoomZ := directionRoomZ + character.Z()
		roomBeforeExploration := roomRepository.FindByXYandZ(targetRoomX, targetRoomY, targetRoomZ)

		result := command.Execute(character, commandDirection.String())

		roomAfterExploration := roomRepository.FindByXYandZ(targetRoomX, targetRoomY, targetRoomZ)
		assert.False(suite.T(), result.HasErrors(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Nil(suite.T(), roomBeforeExploration, fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.NotNil(suite.T(), roomAfterExploration, fmt.Sprintf("Dataset %v %#v", id, dataset))
		biomIsCorrect := assert.Equal(
			suite.T(),
			roomBiom.Air,
			roomAfterExploration.Biom(),
			fmt.Sprintf("Dataset %v %#v", id, dataset),
		)
		if !biomIsCorrect {
			allBiomsAreCorrect = false
		}
		if dataset.canFly {
			assert.Equal(suite.T(), targetRoomX, character.X(), fmt.Sprintf("Dataset %v %#v", id, dataset))
			assert.Equal(suite.T(), targetRoomY, character.Y(), fmt.Sprintf("Dataset %v %#v", id, dataset))
			assert.Equal(suite.T(), targetRoomZ, character.Z(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		} else {
			assert.Equal(suite.T(), initialRoom.X(), character.X(), fmt.Sprintf("Dataset %v %#v", id, dataset))
			assert.Equal(suite.T(), initialRoom.Y(), character.Y(), fmt.Sprintf("Dataset %v %#v", id, dataset))
			assert.Equal(suite.T(), initialRoom.Z(), character.Z(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		}
	}
	if !allBiomsAreCorrect {
		suite.showBiomNumbers(1, len([]roomBiom.Biom{roomBiom.Air}))
	}
}

func (suite *exploreCommandTest) createCharacter() commands.Character {
	return app.Character{}.Create("")
}

func (suite *exploreCommandTest) createRandomWithSeed(seed int64) *random.Random {
	randomizer := random.Random{}.Create()
	randomizer.Seed(seed)

	return randomizer
}

func (suite *exploreCommandTest) provideRoomBioms() []struct {
	randomSeed int64
	biom       roomBiom.Biom
	z          int
} {
	return []struct {
		randomSeed int64
		biom       roomBiom.Biom
		z          int
	}{
		{110, roomBiom.Swamp, 0},
		{113, roomBiom.Hill, 0},
		{89, roomBiom.Water, 0},
		{115, roomBiom.Cave, 0},
		{104, roomBiom.Plain, 0},
		{78, roomBiom.Cliff, 0},
		{114, roomBiom.Sand, 0},
		{79, roomBiom.Town, 0},
		{108, roomBiom.Forest, 0},
		{107, roomBiom.Mountain, 0},
		{176, roomBiom.Swamp, 0},
		{18, roomBiom.Air, 1},
		{8, roomBiom.Mountain, 1},
		{19, roomBiom.Cliff, 1},
	}
}

func (suite *exploreCommandTest) provideDisallowedFromExploreBioms() []struct {
	biom roomBiom.Biom
} {
	return []struct {
		biom roomBiom.Biom
	}{
		{roomBiom.Water},
		{roomBiom.Cliff},
		{roomBiom.Cave},
		{roomBiom.Air},
	}
}

func (suite *exploreCommandTest) provideDisallowedDirections() []struct {
	direction direction.Direction
} {
	return []struct {
		direction direction.Direction
	}{
		{direction.Up},
		{direction.Down},
	}
}

func (suite *exploreCommandTest) provideFlyItem() []struct {
	itemFlag itemFlag.Flag
	canFly   bool
} {
	return []struct {
		itemFlag itemFlag.Flag
		canFly   bool
	}{
		{itemFlag.CanFly, true},
		{itemFlag.MineTool, false},
	}
}

func (suite *exploreCommandTest) showBiomNumbers(z int, biomsCount int) {
	biomSeeds := make(map[roomBiom.Biom]int64)
	var i int64
	for i = 0; len(biomSeeds) < biomsCount; i++ {
		character := suite.createCharacter()
		character.Move(character.X(), character.Y(), z)
		room := app.Room{}.Create(character.X(), character.Y(), character.Z(), roomBiom.Forest)
		roomRepository := app.RoomMemoryRepository{}.Create([]*app.Room{room})
		command := commands.ExploreCommand{}.Create(roomRepository, suite.createRandomWithSeed(i))
		commandDirection := direction.North
		command.Execute(character, commandDirection.String())
		directionX, directionY, directionZ := commandDirection.DiffXYZ()
		targetRoomX := directionX
		targetRoomY := directionY
		targetRoomZ := directionZ + z
		roomAfterExploration := roomRepository.FindByXYandZ(targetRoomX, targetRoomY, targetRoomZ)
		biomSeeds[roomAfterExploration.Biom()] = i
	}

	for biom, iterator := range biomSeeds {
		fmt.Printf("%v %v %v\n", iterator, biom, z)
	}
}

func (suite *exploreCommandTest) assertRoomHasFlags(flags []roomFlag.Flag, room *app.Room, message string) {
	for _, flag := range flags {
		assert.True(suite.T(), room.HasFlag(flag), flag.String(), message)
	}
}
