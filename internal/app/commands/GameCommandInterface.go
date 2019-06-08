package commands

type GameCommand interface {
	Execute(character Character, arguments ...interface{}) (result CommandResult)
}
