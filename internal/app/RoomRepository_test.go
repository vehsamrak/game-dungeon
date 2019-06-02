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

func (suite *roomRepositoryTest) Test_Create_noParameters_roomRepository() {
	repository := app.RoomRepository{}.Create()

	assert.NotNil(suite.T(), repository)
}
