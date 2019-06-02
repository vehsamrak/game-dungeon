package commands

type GameCommand interface {
	Name() string
	Execute(character Character, arguments ...interface{})
}
