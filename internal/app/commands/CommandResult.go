package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
)

type commandResult struct {
	errors map[gameError.Error]bool
}

func (result commandResult) Create() CommandResult {
	return &commandResult{
		errors: make(map[gameError.Error]bool),
	}
}

func (result *commandResult) AddError(err gameError.Error) {
	result.errors[err] = true
}

func (result *commandResult) RemoveError(err gameError.Error) {
	delete(result.errors, err)
}

func (result *commandResult) HasError(err gameError.Error) bool {
	_, ok := result.errors[err]

	return ok
}

func (result *commandResult) HasErrors() bool {
	return len(result.errors) > 0
}

func (result *commandResult) Errors() map[gameError.Error]bool {
	return result.errors
}
