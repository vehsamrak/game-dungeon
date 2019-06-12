package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/random"
)

type Commander struct {
	roomRepository app.RoomRepository
	random         *random.Random
}

func (commander Commander) Create(roomRepository app.RoomRepository, random *random.Random) *Commander {
	return &Commander{roomRepository: roomRepository, random: random}
}

func (commander *Commander) Commands() map[string]GameCommand {
	return map[string]GameCommand{
		"move":    MoveCommand{}.Create(commander.roomRepository),
		"cut":     CutTreeCommand{}.Create(commander.roomRepository),
		"explore": ExploreCommand{}.Create(commander.roomRepository, commander.random),
		"mine":    MineCommand{}.Create(commander.roomRepository, commander.random),
		"fish":    FishCommand{}.Create(commander.roomRepository, commander.random),
	}
}

func (commander *Commander) Command(commandName string) (command GameCommand, err gameError.Error) {
	command, ok := commander.Commands()[commandName]
	if !ok {
		err = gameError.CommandNotFound
	}

	return
}
