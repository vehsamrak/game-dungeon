package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
)

type CommandResult interface {
	LowerHealthOnError() bool
	SetLowerHealthOnError(lowerHealth bool)
	AddError(err gameError.Error)
	HasError(err gameError.Error) bool
	RemoveError(err gameError.Error)
	HasErrors() bool
	Errors() map[gameError.Error]bool
}
