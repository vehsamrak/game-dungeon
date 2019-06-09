package roomFlag

type Flag string

const (
	Road           Flag = "road"
	Unfordable     Flag = "unfordable"
	Trees          Flag = "trees"
	OreProbability Flag = "ore_probability"
)

func (flag Flag) String() string {
	return string(flag)
}
