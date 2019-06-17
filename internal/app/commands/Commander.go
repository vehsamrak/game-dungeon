package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/random"
	"strings"
)

type Commander struct {
	roomRepository app.RoomRepository
	random         *random.Randomizer
}

func (commander Commander) Create(roomRepository app.RoomRepository, random *random.Randomizer) *Commander {
	return &Commander{roomRepository: roomRepository, random: random}
}

func (commander *Commander) Commands() map[string]GameCommand {
	return map[string]GameCommand{
		"move":    MoveCommand{}.Create(commander.roomRepository),
		"cut":     CutTreeCommand{}.Create(commander.roomRepository),
		"explore": ExploreCommand{}.Create(commander.roomRepository, commander.random),
		"mine":    MineCommand{}.Create(commander.roomRepository, commander.random),
		"fish":    FishCommand{}.Create(commander.roomRepository, commander.random),
		"eat":     EatCommand{}.Create(),
	}
}

func (commander *Commander) Execute(character Character, commandWithArguments []string) (
	commandResult CommandResult,
	errors map[gameError.Error]bool,
) {
	commandName := commandWithArguments[0]
	commandArguments := commandWithArguments[1:]
	errors = make(map[gameError.Error]bool)

	command, err := commander.Command(commandName)
	if err != "" {
		errors[err] = true
		return
	}

	characterHealth := character.Health()
	commandHealthPrice := command.HealthPrice()

	if characterHealth <= commandHealthPrice {
		errors[gameError.LowHealth] = true
		return
	}

	commandResult = command.Execute(character, strings.Join(commandArguments, " "))

	if commandResult.HasErrors() {
		for err := range commandResult.Errors() {
			errors[err] = true
		}
	}

	if commandResult.HasErrors() == false || (commandResult.HasErrors() && commandResult.LowerHealthOnError()) {
		character.LowerHealth(commandHealthPrice)
	}

	return
}

func (commander *Commander) Command(commandName string) (command GameCommand, err gameError.Error) {
	command, ok := commander.Commands()[commandName]
	if !ok {
		err = gameError.CommandNotFound
	}

	return
}
