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
	"testing"
)

func TestMoveCommand(test *testing.T) {
	suite.Run(test, &moveCommandTest{})
}

type moveCommandTest struct {
	suite.Suite
}

func (suite *moveCommandTest) Test_Execute_CharacterAndDirectionAndRoomRepository_characterMovedIfRoomExistsAndAvailable() {
	for id, dataset := range suite.provideCharacterDirectionsAndRooms() {
		command := commands.MoveCommand{}.Create(dataset.roomRepository)
		character := suite.getCharacter()
		item := app.Item{}.Create()
		item.AddFlag(dataset.characterItemFlag)
		character.AddItem(item)

		result := command.Execute(character, dataset.direction.String())
		secondResult := command.Execute(character, dataset.direction.String())

		if dataset.error != "" {
			assert.True(suite.T(), result.HasError(dataset.error), fmt.Sprintf("Dataset %v %#v", id, dataset))
		} else {
			assert.False(suite.T(), result.HasErrors(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		}
		assert.Equal(suite.T(), dataset.expectedCharacterX, character.X(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.Equal(suite.T(), dataset.expectedCharacterY, character.Y(), fmt.Sprintf("Dataset %v %#v", id, dataset))
		assert.True(suite.T(), secondResult.HasError(gameError.WaitState), fmt.Sprintf("Dataset %v %#v", id, dataset))
	}
}

func (suite *moveCommandTest) getCharacter() commands.Character {
	character := app.Character{}.Create("")

	return character
}

func (suite *moveCommandTest) provideCharacterDirectionsAndRooms() []struct {
	direction          direction.Direction
	roomRepository     app.RoomRepository
	expectedCharacterX int
	expectedCharacterY int
	expectedCharacterZ int
	error              gameError.Error
	characterItemFlag  itemFlag.Flag
} {
	initialForest := app.Room{}.Create(0, 0, 0, roomBiom.Forest)
	initialWater := app.Room{}.Create(0, 0, 0, roomBiom.Water)
	// upForest := app.Room{}.Create(0, 0, 1, roomBiom.Forest)
	northForest := app.Room{}.Create(0, 1, 0, roomBiom.Forest)
	northWater := app.Room{}.Create(0, 1, 0, roomBiom.Water)
	northAir := app.Room{}.Create(0, 1, 0, roomBiom.Air)

	getRoomRepositoryWithSingleRoom := func(x int, y int, z int, roomFlag roomFlag.Flag) app.RoomRepository {
		room := app.Room{}.Create(x, y, z, roomBiom.Forest)
		room.AddFlag(roomFlag)

		return app.RoomMemoryRepository{}.Create([]*app.Room{initialForest, room})
	}

	getRoomRepositoryWithRooms := func(rooms []*app.Room) app.RoomRepository {
		return app.RoomMemoryRepository{}.Create(rooms)
	}

	getRoomRepositoryWithoutRooms := func() app.RoomRepository {
		return app.RoomMemoryRepository{}.Create(nil)
	}

	return []struct {
		direction          direction.Direction
		roomRepository     app.RoomRepository
		expectedCharacterX int
		expectedCharacterY int
		expectedCharacterZ int
		error              gameError.Error
		characterItemFlag  itemFlag.Flag
	}{
		{direction.North, getRoomRepositoryWithoutRooms(), 0, 0, 0, gameError.RoomNotFound, ""},
		{direction.South, getRoomRepositoryWithRooms([]*app.Room{}), 0, 0, 0, gameError.RoomNotFound, ""},
		{direction.North, getRoomRepositoryWithSingleRoom(0, 1, 0, roomFlag.Road), 0, 1, 0, "", ""},
		{direction.South, getRoomRepositoryWithSingleRoom(0, -1, 0, roomFlag.Road), 0, -1, 0, "", ""},
		{direction.East, getRoomRepositoryWithSingleRoom(1, 0, 0, roomFlag.Road), 1, 0, 0, "", ""},
		{direction.West, getRoomRepositoryWithSingleRoom(-1, 0, 0, roomFlag.Road), -1, 0, 0, "", ""},
		{direction.Up, getRoomRepositoryWithSingleRoom(0, 0, 1, roomFlag.Road), 0, 0, 1, "", ""},
		{direction.Down, getRoomRepositoryWithSingleRoom(0, 0, -1, roomFlag.Road), 0, 0, -1, "", ""},
		{direction.North, getRoomRepositoryWithSingleRoom(0, 1, 0, roomFlag.Unfordable), 0, 0, 0, gameError.RoomUnfordable, ""},
		{direction.North, getRoomRepositoryWithRooms([]*app.Room{initialWater, northWater}), 0, 0, 0, gameError.CantMoveInWater, ""},
		{direction.North, getRoomRepositoryWithRooms([]*app.Room{initialWater, northForest}), 0, 1, 0, "", ""},
		{direction.North, getRoomRepositoryWithRooms([]*app.Room{initialForest, northAir}), 0, 0, 0, gameError.RoomUnfordable, ""},
		{direction.North, getRoomRepositoryWithRooms([]*app.Room{initialForest, northAir}), 0, 1, 0, "", itemFlag.CanFly},
	}
}
