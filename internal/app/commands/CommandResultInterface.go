package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/notice"
)

type CommandResult interface {
	AddError(err gameError.Error)
	HasError(err gameError.Error) bool
	RemoveError(err gameError.Error)
	HasErrors() bool
	AddNotice(notice notice.Notice)
	HasNotice(notice notice.Notice) bool
}
