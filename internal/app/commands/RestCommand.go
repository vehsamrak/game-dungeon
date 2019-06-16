package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"time"
)

type RestCommand struct {
	roomRepository app.RoomRepository
	waitState      time.Duration
	healthPrice    int
	healthGain     int
}

func (command *RestCommand) HealthPrice() int {
	return command.healthPrice
}

func (command RestCommand) Create() *RestCommand {
	return &RestCommand{
		waitState:  10 * time.Second,
		healthGain: 2,
	}
}

func (command *RestCommand) Execute(character Character, arguments ...string) (result CommandResult) {
	result = commandResult{}.Create()

	if character.HasActiveTimers() {
		result.AddError(gameError.WaitState)
		return
	}

	if character.FullHealth() {
		result.AddError(gameError.HealthFull)
	} else if character.Health()+command.healthGain > character.MaxHealth() {
		character.RestoreHealth()
	} else {
		character.IncreaseHealth(command.healthGain)
	}

	return
}
