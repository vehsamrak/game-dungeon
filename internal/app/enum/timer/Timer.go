package timer

type Timer string

const (
	Rest           Timer = "rest and move command"
	Explore        Timer = "explore command"
	GatherResource Timer = "gather any resource"
)
