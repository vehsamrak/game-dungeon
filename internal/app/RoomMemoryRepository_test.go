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

func (suite *roomRepositoryTest) Test_FindByXYandZ_existingXYZ_room() {
	repository := app.RoomMemoryRepository{}.Create(nil)
	x, y, z := 0, 0, 0

	room := repository.FindByXYandZ(x, y, z)

	assert.NotNil(suite.T(), room)
	assert.Equal(suite.T(), x, room.X())
	assert.Equal(suite.T(), y, room.Y())
	assert.Equal(suite.T(), z, room.Z())
}

func (suite *roomRepositoryTest) Test_FindByXYandZ_nonexistentXYZ_nil() {
	repository := app.RoomMemoryRepository{}
	nonexistentX := 0
	nonexistentY := 0
	nonexistentZ := 0

	room := repository.FindByXYandZ(nonexistentX, nonexistentY, nonexistentZ)

	assert.Nil(suite.T(), room)
}

func (suite *roomRepositoryTest) Test_FindByXY_characterWithExistingXY_room() {
	x, y, z := 1, 1, 0
	repository := app.RoomMemoryRepository{}
	repository.AddRoom(app.Room{}.Create(x, y, z, roomBiom.Forest))
	character := &app.Character{}
	character.Move(x, y, z)

	room := repository.FindByXYZ(character)

	assert.NotNil(suite.T(), room)
	assert.Equal(suite.T(), character.X(), room.X())
	assert.Equal(suite.T(), character.Y(), room.Y())
}

func (suite *roomRepositoryTest) Test_FindByXY_characterWithNonexistentXY_room() {
	repository := app.RoomMemoryRepository{}
	character := &app.Character{}

	room := repository.FindByXYZ(character)

	assert.Nil(suite.T(), room)
}

func (suite *roomRepositoryTest) Test_AddRoom_newRoomCreated_newRoomAddedToRepository() {
	x, y, z := 0, 0, 0
	repository := app.RoomMemoryRepository{}
	nonexistentRoom := repository.FindByXYandZ(x, y, z)
	room := app.Room{}.Create(x, y, z, roomBiom.Forest)

	repository.AddRoom(room)

	existingRoom := repository.FindByXYandZ(x, y, z)
	assert.Nil(suite.T(), nonexistentRoom)
	assert.NotNil(suite.T(), existingRoom)
}
