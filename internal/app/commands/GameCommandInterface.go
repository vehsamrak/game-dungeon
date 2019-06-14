package commands

type GameCommand interface {
	Execute(character Character, arguments ...string) (result CommandResult)
	HealthPrice() int
}
