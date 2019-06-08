package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app/enum/notice"
)

type CommandResult interface {
	// todo pass exception enum
	AddError(err error)
	HasError(err error) bool
	HasErrors() bool
	AddNotice(notice notice.Notice)
	HasNotice(notice notice.Notice) bool
	RemoveError(err error)
}
