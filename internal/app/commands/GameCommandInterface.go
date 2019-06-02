package commands

type GameCommand interface {
	Execute(character Character, arguments ...interface{}) (err error)
}
