package timer

type Timer string

const (
	Move           Timer = "move command"
	Explore        Timer = "explore command"
	GatherResource Timer = "gather any resource"
)
