package app_test

import (
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomBiom"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
)

func TestRoomRepository(test *testing.T) {
	suite.Run(test, &roomRepositoryTest{})
}

type roomRepositoryTest struct {
	suite.Suite
}

func (suite *roomRepositoryTest) Test_FindByXandY_existingXY_room() {
	repository := app.RoomMemoryRepository{}.Create(nil)
	x, y := 1, 1

	room := repository.FindByXandY(x, y)

	assert.NotNil(suite.T(), room)
	assert.Equal(suite.T(), x, room.X())
	assert.Equal(suite.T(), y, room.Y())
}

func (suite *roomRepositoryTest) Test_FindByXandY_nonexistentXY_nil() {
	repository := app.RoomMemoryRepository{}
	nonexistentX := 0
	nonexistentY := 0

	room := repository.FindByXandY(nonexistentX, nonexistentY)

	assert.Nil(suite.T(), room)
}

func (suite *roomRepositoryTest) Test_FindByXY_characterWithExistingXY_room() {
	x, y := 1, 1
	repository := app.RoomMemoryRepository{}
	repository.AddRoom(app.Room{}.Create(x, y, roomBiom.Forest))
	character := &app.Character{}
	character.Move(x, y)

	room := repository.FindByXY(character)

	assert.NotNil(suite.T(), room)
	assert.Equal(suite.T(), character.X(), room.X())
	assert.Equal(suite.T(), character.Y(), room.Y())
}

func (suite *roomRepositoryTest) Test_FindByXY_characterWithNonexistentXY_room() {
	repository := app.RoomMemoryRepository{}
	character := &app.Character{}

	room := repository.FindByXY(character)

	assert.Nil(suite.T(), room)
}

func (suite *roomRepositoryTest) Test_AddRoom_newRoomCreated_newRoomAddedToRepository() {
	x, y := 0, 0
	repository := app.RoomMemoryRepository{}
	nonexistentRoom := repository.FindByXandY(x, y)
	room := app.Room{}.Create(x, y, roomBiom.Forest)

	repository.AddRoom(room)

	existingRoom := repository.FindByXandY(x, y)
	assert.Nil(suite.T(), nonexistentRoom)
	assert.NotNil(suite.T(), existingRoom)
}
