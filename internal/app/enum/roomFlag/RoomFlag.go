package roomFlag

type Flag string

const (
	Road       Flag = "road"
	Unfordable Flag = "unfordable"
	Trees      Flag = "trees"
)

func ActiveFlags() []Flag {
	return []Flag{
		Road,
		Unfordable,
		Trees,
	}
}
