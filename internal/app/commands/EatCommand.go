package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/itemFlag"
)

type EatCommand struct {
	healthPrice int
	healthGain  int
}

func (command *EatCommand) HealthPrice() int {
	return command.healthPrice
}

func (command EatCommand) Create() *EatCommand {
	return &EatCommand{healthGain: 10}
}

func (command *EatCommand) Execute(character Character, arguments ...string) (result CommandResult) {
	result = commandResult{}.Create()

	if character.HasItemFlag(itemFlag.Food) {
		item := character.FindItemWithFlag(itemFlag.Food)
		character.DropItem(item)
		foodHealthGain := command.healthGain
		if character.Health()+foodHealthGain > character.MaxHealth() {
			character.RestoreHealth()
		} else {
			character.IncreaseHealth(foodHealthGain)
		}
	} else {
		result.AddError(gameError.FoodNotFound)
	}

	return
}
