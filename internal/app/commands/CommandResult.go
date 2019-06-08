package commands

import "github.com/vehsamrak/game-dungeon/internal/app/enum/notice"

type commandResult struct {
	errors  map[error]bool
	notices map[notice.Notice]bool
}

func (result commandResult) Create() CommandResult {
	return &commandResult{
		errors:  make(map[error]bool),
		notices: make(map[notice.Notice]bool),
	}
}

func (result *commandResult) AddError(err error) {
	result.errors[err] = true
}

func (result *commandResult) HasError(err error) bool {
	_, ok := result.errors[err]

	return ok
}

func (result *commandResult) HasErrors() bool {
	return len(result.errors) > 0
}

func (result *commandResult) HasNotice(notice notice.Notice) bool {
	_, ok := result.notices[notice]

	return ok
}

func (result *commandResult) AddNotice(notice notice.Notice) {
	result.notices[notice] = true
}
