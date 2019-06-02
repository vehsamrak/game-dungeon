package app_test

import (
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

func (suite *roomRepositoryTest) Test_FindByXY_existingXY_room() {
	repository := app.RoomMemoryRepository{}.Create(nil)
	x, y := 1, 1

	room := repository.FindByXY(x, y)

	assert.NotNil(suite.T(), room)
	assert.Equal(suite.T(), x, room.X())
	assert.Equal(suite.T(), y, room.Y())
}

func (suite *roomRepositoryTest) Test_FindByXY_nonexistentXY_nil() {
	repository := app.RoomMemoryRepository{}
	nonexistentX := 0
	nonexistentY := 0

	room := repository.FindByXY(nonexistentX, nonexistentY)

	assert.Nil(suite.T(), room)
}
