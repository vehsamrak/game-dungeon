package app_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomBiom"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
	"testing"
)

func TestRoom(test *testing.T) {
	suite.Run(test, &roomTest{})
}

type roomTest struct {
	suite.Suite
}

func (suite *roomTest) Test_Create_roomParameters_newRoomCreated() {
	x, y := 1, 1

	room := app.Room{}.Create(x, y, roomBiom.Forest)

	assert.Equal(suite.T(), x, room.X())
	assert.Equal(suite.T(), y, room.Y())
}

func (suite *roomTest) Test_AddFlag_roomWithoutFlags_flagsAddedToRoom() {
	room := suite.createRoom()
	firstFlag := roomFlag.Trees
	secondFlag := roomFlag.Road
	flags := []roomFlag.Flag{secondFlag}

	room.AddFlag(firstFlag)
	room.AddFlags(flags)

	assert.True(suite.T(), room.HasFlag(firstFlag))
	assert.True(suite.T(), room.HasFlag(secondFlag))
}

func (suite *roomTest) Test_Flags_roomWithFlag_flagReturned() {
	room := suite.createRoom()
	room.AddFlag(roomFlag.Unfordable)

	flags := room.Flags()

	assert.Len(suite.T(), flags, 1)
}

func (suite *roomTest) createRoom() *app.Room {
	return app.Room{}.Create(0, 0, roomBiom.Forest)
}
